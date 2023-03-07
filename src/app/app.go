package app

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"git.unemeta.com/Backstage/une/src/config"
	"git.unemeta.com/Backstage/une/src/env"
	"git.unemeta.com/Backstage/une/src/util"
	migrate "github.com/rubenv/sql-migrate"

	mysql "github.com/go-sql-driver/mysql"
)

type App interface {
	Api(n string)
	Rpc()
	Run(module string)
	Doc(enum config.DocEnum)
	MigrateUp(env config.SqlEnvEnum)
	MigrateDown(env config.SqlEnvEnum)
	MigrateNew(name string, open bool)
	MigrateModel(env config.SqlEnvEnum, useCache bool)
	Name() string
	Path() string
}

type defaultApp struct {
	name string
	path string
}

func (d defaultApp) Run(module string) {
	runCmdStr := fmt.Sprintf("cd %s/%s && go run %s.go", d.Path(), module, d.Name())
	util.MustCmdSpawn(runCmdStr)
}

func (d defaultApp) Api(n string) {
	apiCmdStr := fmt.Sprintf("goctl api go -api %s/%s/desc/%s.api -dir %s/%s --home %s --style goZero",
		d.Path(), n, getApiName(d.Path(), d.Name(), n), d.Path(), n, templateDir)
	util.MustCmd(apiCmdStr)
}

func (d defaultApp) Rpc() {
	rpcCmdStr := fmt.Sprintf("cd %s/rpc && goctl rpc protoc pb/%s.proto --go_out=. --go-grpc_out=. --zrpc_out=. --home=%s --style=goZero",
		d.Path(), d.Name(), templateDir)
	util.MustCmd(rpcCmdStr)
}

func (d defaultApp) Doc(enum config.DocEnum) {
	env.MustInitDocEnv()
	docDir := path.Join(config.GetConfig().ProjectDir, "deploy", "doc")
	docCmdStr := fmt.Sprintf("goctl api plugin -plugin goctl-swagger='swagger -filename %s.json -host api.test.unemeta.com -%s' -api %s/api/desc/%s.api -dir '%s'",
		d.Name(), "https", d.Path(), getApiName(d.Path(), d.Name(), "api"), docDir)
	util.MustCmd(docCmdStr)
	switch enum {
	case config.DocAll:
		panic("not implement")
	case config.DocYApi:
		yApiCmdStr := fmt.Sprintf("npx -p yapi-cli yapi import --config %s", config.GetYApiConfigFileStr(d.Name()))
		util.MustCmdSpawn(yApiCmdStr)
	case config.DocReadMe:
		id := config.GetReadMeConfigStr(d.Name())
		appDocFile := path.Join(config.GetConfig().ProjectDir, "deploy", "doc", fmt.Sprintf("%s.json", d.Name()))
		if id != "" {
			readMeCmd := fmt.Sprintf(`npx -p rdme@latest rdme openapi %s --key=%s --id=%s`,
				appDocFile, config.GetConfig().ReadMe.Key, id)
			util.MustCmdSpawn(readMeCmd)
		}
	default:
		break
	}
}

func (d defaultApp) MigrateUp(env config.SqlEnvEnum) {
	var wg sync.WaitGroup
	for _, temp := range d.dsn(env) {
		wg.Add(1)
		sqlUri := temp
		go func() {
			migrateWrap(sqlUri, d.Name(), func(db *sql.DB, migrations *migrate.FileMigrationSource) {
				n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
				if err != nil {
					fmt.Printf("migrate up err:%+v\n", err)
					return
				}
				fmt.Printf("Applied %d migrations up!\n", n)
			})
			wg.Done()
		}()
	}
	wg.Wait()
}

func (d defaultApp) MigrateModel(env config.SqlEnvEnum, useCache bool) {
	modelDir := getModelDir(d.Path())
	if modelDir != "" {
		conf, _ := mysql.ParseDSN(d.dsn(env)[0])
		goctlDbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.User, conf.Passwd, conf.Addr, conf.DBName)
		var updateModelCmd string
		if useCache {
			updateModelCmd = fmt.Sprintf("cd %s && goctl model mysql datasource -url='%s' --home %s -table='*' -c -dir . && rm migrationsmodel.go &&rm migrationsmodel_gen.go",
				modelDir, goctlDbUri, templateDir)
		} else {
			updateModelCmd = fmt.Sprintf("cd %s && goctl model mysql datasource -url='%s' --home %s -table='*' -dir . && rm migrationsmodel.go &&rm migrationsmodel_gen.go",
				modelDir, goctlDbUri, templateDir)
		}
		util.MustCmd(updateModelCmd)
	} else {
		fmt.Printf("generate goctl model failed, model not found in %s/model or %s/rpc/dir\n", d.Path(), d.Path())
	}
}

func (d defaultApp) MigrateDown(env config.SqlEnvEnum) {
	var wg sync.WaitGroup
	for _, temp := range d.dsn(env) {
		wg.Add(1)
		sqlUri := temp
		go func() {
			migrateWrap(sqlUri, d.Name(), func(db *sql.DB, migrations *migrate.FileMigrationSource) {
				n, err := migrate.ExecMax(db, "mysql", migrations, migrate.Down, 1)
				if err != nil {
					fmt.Printf("migrate down err:%+v\n", err)
					return
				}
				fmt.Printf("Applied %d migration down!\n", n)
			})
			wg.Done()
		}()
	}
	wg.Wait()
}

const migrateNewTimeLayout = "20060102150405"

func (d defaultApp) MigrateNew(name string, open bool) {
	fileName := fmt.Sprintf("%s-%s.sql", time.Now().Format(migrateNewTimeLayout), name)
	migrateFile := path.Join(config.GetConfig().ProjectDir, "deploy", "sql", d.Name(), fileName)
	file, err := os.Create(migrateFile)
	if err != nil {
		fmt.Printf("create file:%s err:%+v\n", migrateFile, err)
		return
	}
	_, err = file.WriteString(`-- +migrate Up

-- +migrate Down
`)
	if err != nil {
		fmt.Printf("write template to file:%s err:%+v\n", migrateFile, err)
		return
	}
	err = file.Close()
	if err != nil {
		fmt.Printf("close file:%s err:%+v\n", migrateFile, err)
		return
	}
	if open {
		if !util.AppExist("datagrip") {
			fmt.Println("datagrip not found in path.")
			return
		}
		util.MustCmd(fmt.Sprintf("datagrip %s", migrateFile))
	}
}

func (d defaultApp) Name() string {
	return d.name
}

func (d defaultApp) Path() string {
	return d.path
}

func (d defaultApp) dsn(e config.SqlEnvEnum) []string {
	var res []string
	var modules []string
	var common []string
	switch e {
	case config.DbLocal:
		modules = config.GetConfig().Db.Local.Modules
		common = util.GetValueFromBracket(config.GetConfig().Db.Local.Common)
	case config.DbTest:
		modules = config.GetConfig().Db.Test.Modules
		common = util.GetValueFromBracket(config.GetConfig().Db.Test.Common)
	}
	for _, m := range modules {
		dbSourceStr := util.GetValueFromBracket(m)
		if dbSourceStr[0] == d.Name() {
			dsnStr := fmt.Sprintf("%s@tcp(%s)/%s?parseTime=true", dbSourceStr[2], dbSourceStr[3], dbSourceStr[1])
			res = append(res, dsnStr)
		}
	}
	if len(res) < 1 {
		dsnStr := fmt.Sprintf("%s@tcp(%s)/unemeta_%s?parseTime=true", common[0], common[1], d.Name())
		res = append(res, dsnStr)
	}
	if len(res) < 1 {
		panic(fmt.Sprintf("sql source not found in this module:%s", d.Name()))
	}
	return res
}

func migrateWrap(sqlUri string, name string, callback func(dataSource *sql.DB, migrations *migrate.FileMigrationSource)) {
	migrations := &migrate.FileMigrationSource{
		Dir: path.Join(config.GetConfig().ProjectDir, "deploy", "sql", name),
	}
	fmt.Printf("%s\n", sqlUri)
	db, err := sql.Open("mysql", sqlUri)
	migrate.SetTable("migrations")
	if err != nil {
		fmt.Printf("open mysql uri:%s err:%+v\n", sqlUri, err)
		return
	}
	callback(db, migrations)
}

func getModelDir(appPath string) string {
	if util.AssetExist(path.Join(appPath, "model")) {
		return path.Join(appPath, "model")
	}
	if util.AssetExist(path.Join(appPath, "rpc", "model")) {
		return path.Join(appPath, "rpc", "model")
	}
	return ""
}

// return root if root.api exist.
func getApiName(p string, name string, module string) string {
	if util.AssetExist(path.Join(p, module, "desc", "root.api")) {
		return "root"
	} else {
		return name
	}
}
