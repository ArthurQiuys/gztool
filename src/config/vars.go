package config

import (
	"fmt"
	"io"
	"os"
	"path"
	"sync"

	"git.unemeta.com/Backstage/une/src/config/compatibility"
	"git.unemeta.com/Backstage/une/src/util"
	"github.com/naoina/toml"
)

const (
	dbSourceCommon = "dbSourceCommon"
)

const YApiConfigTemplate = `{
  "type": "swagger",
  "token": "%s",
  "file": "%s",
  "merge": "mergin",
  "server": "%s"
}
`

var configPath = fmt.Sprintf("%s/.unemeta_cli.toml", util.HomeDir())

var config compatibility.Config
var configOnce sync.Once

func getConfigOnce() compatibility.Config {
	configOnce.Do(func() {
		if !util.AssetExist(configPath) {
			fmt.Printf("config file:%s not found, may be you should update or generate default config.\n", configPath)
			fmt.Printf("[y/yes] try migrate by config compatibility layers.\n")
			fmt.Printf("[n/no] to generate a default config\n")
			var input string
			for {
				_, err := fmt.Scanln(&input)
				if err != nil && err.Error() == "unexpected newline" {
					continue
				} else if input != "y" && input != "yes" && input != "n" && input != "no" {
					fmt.Printf("Please input [y/yes] or [n/no]\n")
					continue
				}
				break
			}
			switch input {
			case "y", "yes":
				UpdateConfig()
			case "n", "no":
				config = *defaultConfig()
				data, _ := toml.Marshal(config)
				configFile, err := os.Create(configPath)
				if err != nil {
					panic(fmt.Sprintf("create config file err:%+v path:%s", err, configPath))
				}
				_, err = configFile.WriteString(string(data))
				defer configFile.Close()
				if err != nil {
					panic(fmt.Sprintf("write config to file %s err:%+v", configPath, err))
				}
			}
		} else {
			configFile, err := os.Open(configPath)
			if err != nil {
				panic(fmt.Sprintf("open config file err:%+v path:%s", err, configPath))
			}
			data, err := io.ReadAll(configFile)
			if err != nil {
				panic(fmt.Sprintf("read from config file err:%+v path:%s", err, configPath))
			}
			toml.Unmarshal(data, &config)
			if err != nil {
				panic(fmt.Sprintf("unmarshal config file err:%+v path:%s", err, configPath))
			}
			configFile.Close()
		}
		if !util.AssetExist(config.CacheDir) {
			if config.CacheDir == "" {
				panic(fmt.Sprintf("cache dir is empty"))
			}
			err := os.MkdirAll(config.CacheDir, 0751)
			if err != nil {
				panic(fmt.Sprintf("create cache dir err:%+v path:%s", err, config.CacheDir))
			}
		}
		if !util.AssetExist(path.Join(config.CacheDir, "bin")) {
			err := os.MkdirAll(path.Join(config.CacheDir, "bin"), 0751)
			if err != nil {
				panic(fmt.Sprintf("create cache bin dir err:%+v path:%s", err, path.Join(config.CacheDir, "bin")))
			}
		}
	})
	return config
}
