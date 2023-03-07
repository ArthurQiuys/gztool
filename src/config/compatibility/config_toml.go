package compatibility

import (
	"fmt"
	"io"
	"os"

	"git.unemeta.com/Backstage/une/src/util"
	"github.com/naoina/toml"
	"gopkg.in/yaml.v3"
)

type TomlConfig struct {
	oldPath string
	newPath string
}

func NewTomlConfig(oldPath string, newPath string) *TomlConfig {
	return &TomlConfig{
		oldPath: oldPath,
		newPath: newPath,
	}
}

func (t TomlConfig) Compatility(config *Config) {
	if !util.AssetExist(t.oldPath) {
		return
	}
	errorStr := "migrate config from yaml to toml error"
	configFile, err := os.Open(t.oldPath)
	if err != nil {
		panic(fmt.Sprintf("%s, open config file err:%+v path:%s", errorStr, err, t.oldPath))
	}
	data, err := io.ReadAll(configFile)
	if err != nil {
		panic(fmt.Sprintf("%s, read from config file err:%+v path:%s", errorStr, err, t.oldPath))
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(fmt.Sprintf("%s, unmarshal config file err:%+v path:%s", errorStr, err, t.oldPath))
	}
	configFile.Close()
	newConfigStr, err := toml.Marshal(config)
	if err != nil {
		panic(fmt.Sprintf("%s, convert config to toml err:%+v", errorStr, err))
	}
	configFile, err = os.Create(t.newPath)
	if err != nil {
		panic(fmt.Sprintf("%s, create config file err:%+v path:%s", errorStr, err, t.newPath))
	}
	_, err = configFile.WriteString(string(newConfigStr))
	if err != nil {
		panic(fmt.Sprintf("%s, write config to file err:%+v path:%s", errorStr, err, t.newPath))
	}
}
