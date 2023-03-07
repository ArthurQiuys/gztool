package app

import (
	"path"
	"strings"
	"sync"

	"git.unemeta.com/Backstage/une/src/env"
	"git.unemeta.com/Backstage/une/src/util"

	"git.unemeta.com/Backstage/une/src/config"
)

var templateDir string

type Apps struct {
	apps  []App
	names []string
	paths []string
}

var apps Apps
var appOnce sync.Once

func GetApp(name string) App {
	for _, app := range getApps().apps {
		if strings.HasPrefix(app.Name(), name) {
			return app
		}
	}
	return nil
}

func GetNames() []string {
	return getApps().names
}

func GetPaths() []string {
	return getApps().names
}

func GetAllApp() []App {
	return getApps().apps
}

func getApps() Apps {
	appOnce.Do(func() {
		env.MustInitEnv()
		appDir := path.Join(config.GetConfig().ProjectDir, "app")
		appDirs := util.GetSubDirectories(appDir)
		for _, name := range appDirs {
			apps.apps = append(apps.apps, defaultApp{name: name, path: path.Join(appDir, name)})
			apps.names = append(apps.names, name)
			apps.paths = append(apps.paths, path.Join(appDir, name))
		}
		templateDir = path.Join(config.GetConfig().CacheDir, "goctl_template")
	})
	return apps
}
