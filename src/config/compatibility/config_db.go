package compatibility

type DbDocConfigCompatility struct {
	defaultConfig *Config
}

func NewDbDocConfigCompatility(defaultConfig *Config) *DbDocConfigCompatility {
	return &DbDocConfigCompatility{
		defaultConfig: defaultConfig,
	}
}

func (d DbDocConfigCompatility) Compatility(config *Config) {
	if len(config.Db.Local.Modules) < 0 {
		config.Db.Local.Modules = d.defaultConfig.Db.Local.Modules
	}
	if config.Db.Local.Common == "" {
		config.Db.Local.Common = d.defaultConfig.Db.Local.Common

	}
	if len(config.Db.Test.Modules) < 1 {
		config.Db.Test.Modules = d.defaultConfig.Db.Test.Modules
	}
	if len(config.YApi.Modules) < 1 {
		config.YApi.Modules = d.defaultConfig.YApi.Modules
	}
	if config.YApi.Server == "" {
		config.YApi.Server = d.defaultConfig.YApi.Server
	}
	if len(config.ReadMe.Modules) < 1 {
		config.ReadMe.Modules = d.defaultConfig.ReadMe.Modules
	}
	if config.ReadMe.Key == "" {
		config.ReadMe.Key = d.defaultConfig.ReadMe.Key
	}
}
