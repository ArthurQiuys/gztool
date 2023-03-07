package env

import (
	"fmt"
	"log"
	"os"
	"path"

	"git.unemeta.com/Backstage/une/src/config"
	"git.unemeta.com/Backstage/une/src/util"
)

func MustInitDocEnv() {
	apiDir := path.Join(config.GetConfig().CacheDir, "api")
	if !util.AssetExist(apiDir) {
		err := os.MkdirAll(apiDir, 0751)
		if err != nil {
			panic(fmt.Sprintf("create cache dir err:%+v path:%s", err, apiDir))
		}
	}
	// 检查yApi配置文件
	for _, module := range config.GetConfig().YApi.Modules {
		yApiConf := util.GetValueFromBracket(module)
		if len(yApiConf) < 2 {
			continue
		}
		appName := yApiConf[0]
		yApiConfFileStr := config.GetYApiConfigFileStr(appName)
		if !util.AssetExist(yApiConfFileStr) {
			log.Printf("yApi config file:%s not exist, while create it.", yApiConfFileStr)
			appDocFile := path.Join(config.GetConfig().ProjectDir, "deploy", "doc", fmt.Sprintf("%s.json", appName))
			yApiConfStr := fmt.Sprintf(config.YApiConfigTemplate, yApiConf[1], appDocFile, config.GetConfig().YApi.Server)
			file, err := os.Create(yApiConfFileStr)
			if err != nil {
				log.Printf("create yapi config file:%s err:%+v", yApiConfFileStr, err)
				continue
			}
			_, err = file.WriteString(yApiConfStr)
			if err != nil {
				log.Printf("write yapi config to file:%s err:%+v", yApiConfFileStr, err)
				continue
			}
			file.Close()
		}
	}
}
