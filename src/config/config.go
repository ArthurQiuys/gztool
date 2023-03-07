package config

import (
	"fmt"
	"os"

	"git.unemeta.com/Backstage/une/src/config/compatibility"
	"git.unemeta.com/Backstage/une/src/util"
)

func defaultConfig() *compatibility.Config {
	cacheDir := fmt.Sprintf("%s/.cache/unemeta_cli/", util.HomeDir())
	pwdDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("get current dir err:%+v", err))
	}
	return &compatibility.Config{
		ProjectDir: pwdDir,
		CacheDir:   cacheDir,
		Db: compatibility.DbConfig{
			Local: compatibility.DbEnvConfig{
				Common: "{{unemeta_backend_dev:uneune202}}{{localhost:3069}}",
				Modules: []string{
					"{{user}}{{unemeta_users}}{{unemeta_backend_dev:uneune202}}{{localhost:3069}}",
				},
			},
			Test: compatibility.DbEnvConfig{
				Common: "{{unemeta_backend_dev:UneUne202.}}{{34.85.98.88:3306}}",
				Modules: []string{
					"{{console}}{{unemeta_console}}{{unemeta_admin:UneUne202.}}{{34.85.98.88:3306}}",
					"{{user}}{{unemeta_users}}{{unemeta_backend_dev:UneUne202.}}{{34.85.98.88:3306}}",
					"{{market}}{{unemeta_market}}{{unemeta_backend_dev:UneUne202.}}{{34.85.98.88:3306}}",
					"{{market}}{{unemeta_market}}{{unemeta_backend_dev:UneUne202.}}{{34.84.252.50:3306}}",
				},
			},
		},
		YApi: compatibility.YApiConfig{
			Server: "http://35.209.215.76:40001",
			Modules: []string{
				"{{console}}{{d37b8b56ed439017efd4bcbe08f1df60d84f2ff127dac751d4f6d88fd27952ad}}",
				"{{user}}{{0eda057be16656f12676d1f00f34ded8eaedfcfe0e947ca45cd5d97a69b86a27}}",
				"{{integral}}{{ff2d5df9deac8567259aa978e9c534a6e71038e12a606d7635b1d294eda26b1b}}",
				"{{systeminfo}}{{74146e49063c48fe6cb2a5208aa1292f18d4651a21d55b944fbfba1eb6d5323c}}",
				"{{market}}{{968467dca328d146a1bdf06ad8daf098d9a15f5cc48d38fbad0e90aad6f7be58}}",
				"{{transaction}}{{c9e1d472f549da98d2368332d1b6fe650e2225f17b4095912b612856d60d94cc}}",
			},
		},
		//rdme openapi console.json \
		//  --key=rdme_xn8s9h17429f437a5a8a97371d49cab3b9a458698f43f045abf7b46601fe0195ee3c30 \
		//--id=639c8750455b32000f346c98
		ReadMe: compatibility.ReadMeConfig{
			Key: "rdme_xn8s9h17429f437a5a8a97371d49cab3b9a458698f43f045abf7b46601fe0195ee3c30",
			Modules: []string{
				"{{console}}{{639c8750455b32000f346c98}}",
				"{{integral}}{{639c895f9fb014026fe471c2}}",
				"{{market}}{{639c8981178d5606e01c47cc}}",
				"{{systeminfo}}{{639c89a9ba635200467f0e5d}}",
				"{{transaction}}{{639c89c70879c50093fb1b8d}}",
				"{{user}}{{639c89e01f22560012be3fce}}",
			},
		},
	}
}
