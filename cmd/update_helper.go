package cmd

import (
	"fmt"
	"git.unemeta.com/Backstage/une/src/config"
	"git.unemeta.com/Backstage/une/src/util"
	"path"
)

func UpdateGoCtlSwagger() {
	binDir := path.Join(config.GetConfig().CacheDir, "bin")
	goSwaggerBin := path.Join(binDir, "goctl-swagger")
	goSwaggerCache := path.Join(config.GetConfig().CacheDir, "goctl-swagger")
	if util.AssetExist(goSwaggerCache) {
		util.MustCmd(fmt.Sprintf(`cd '%s' && \
git pull origin main && \
go build -o %s main.go`, goSwaggerCache, goSwaggerBin))
	} else {
		util.MustCmd(fmt.Sprintf(`git clone --depth=1 %s %s && \
go build -o %s main.go`, config.GoCtlSwaggerGit, goSwaggerCache, goSwaggerBin))
	}
}

func UpdateGoCtlTemplate() {
	templateDir := path.Join(config.GetConfig().CacheDir, "goctl_template")
	util.MustCmd(fmt.Sprintf(`cd %s && \
git pull origin master`, templateDir))
}
