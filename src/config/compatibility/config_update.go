package compatibility

import (
	"fmt"
	"os"
)

type ConfigUpdate struct {
	updatePath  string
	marshalFunc func(any) ([]byte, error)
}

func NewConfigUpdate(updatePath string, marshalFunc func(any) ([]byte, error)) *ConfigUpdate {
	return &ConfigUpdate{
		updatePath:  updatePath,
		marshalFunc: marshalFunc,
	}
}

func (c ConfigUpdate) Compatility(config *Config) {
	data, _ := c.marshalFunc(config)
	configFile, err := os.OpenFile(c.updatePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0751)
	if err != nil {
		panic(fmt.Sprintf("open config file %s err:%+v", c.updatePath, err))
	}
	_, err = configFile.WriteString(string(data))
	defer configFile.Close()
	if err != nil {
		panic(fmt.Sprintf("write config file %s err:%+v", c.updatePath, err))
	}
}
