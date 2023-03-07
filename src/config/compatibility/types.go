package compatibility

type Config struct {
	ProjectDir string       `yaml:"project_dir"`
	CacheDir   string       `yaml:"cache_dir"`
	Db         DbConfig     `yaml:"db_config"`
	YApi       YApiConfig   `yaml:"y_api"`
	ReadMe     ReadMeConfig `yaml:"read_me"`
}

type DbConfig struct {
	Local DbEnvConfig `yaml:"local"`
	Test  DbEnvConfig `yaml:"test"`
}

type DbEnvConfig struct {
	Common  string   `yaml:"common"`  // {{user}}{{passwd}}
	Modules []string `yaml:"modules"` // {{appName}}{{dbName}}{{user}}{{passwd}}
}

type YApiConfig struct {
	Server  string   `yaml:"server"`
	Modules []string `yaml:"modules"` // {{appName}}{{token}}
}

type ReadMeConfig struct {
	Key     string   `yaml:"key"`
	Modules []string `yaml:"modules"`
}

type ConfigCompatibility interface {
	Compatility(config *Config)
}
