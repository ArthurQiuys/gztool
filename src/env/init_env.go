package env

import (
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"git.unemeta.com/Backstage/une/src/config"
	"git.unemeta.com/Backstage/une/src/util"
)

var initEnvOnce sync.Once

func MustInitEnv() {
	initEnvOnce.Do(func() {
		binDir := path.Join(config.GetConfig().CacheDir, "bin")
		pathEnv := os.Getenv("PATH")
		_ = os.Setenv("PATH", fmt.Sprintf("%s:%s", binDir, pathEnv))
		templateDir := path.Join(config.GetConfig().CacheDir, "goctl_template")
		if !util.AssetExist(templateDir) {
			log.Printf("goctl_template not exist, clone to cache dir:%s", templateDir)
			util.MustCmd(fmt.Sprintf("git clone --depth=1 %s %s", config.GoCtlTemplateGit, templateDir))
		}
		if !util.AppExist("goctl") {
			log.Println("goctl not found, go install goctl")
			util.MustCmd("go install " + config.GoCtlInstall)
		}

		goSwaggerBin := path.Join(binDir, "goctl-swagger")
		if !util.AssetExist(goSwaggerBin) {
			goSwaggerCache := path.Join(config.GetConfig().CacheDir, "goctl-swagger")
			if !util.AssetExist(goSwaggerCache) {
				log.Printf("goctl-swagger not exist, clone to cache dir:%s", goSwaggerCache)
				util.MustCmd(fmt.Sprintf("git clone --depth=1 %s %s", config.GoCtlSwaggerGit, goSwaggerCache))
			}
			log.Printf("goctl-swagger bin not exist, build to bin dir:%s", goSwaggerBin)
			util.MustCmd(fmt.Sprintf("cd %s && go build -o %s main.go", goSwaggerCache, goSwaggerBin))
		}
		if !util.AppExist("sql-migrate") {
			log.Println("sql-migrate not fount, go install sql-migrate")
			util.MustCmd("go install github.com/rubenv/sql-migrate/...@latest")
		}
	})
}
