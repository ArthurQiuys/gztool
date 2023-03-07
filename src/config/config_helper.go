package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"git.unemeta.com/Backstage/une/src/config/compatibility"
	"git.unemeta.com/Backstage/une/src/util"
	"github.com/naoina/toml"
	"github.com/spf13/cobra"
)

func GetYApiConfigFileStr(name string) string {
	return path.Join(GetConfig().CacheDir, "api", fmt.Sprintf("yApi-%s.json", name))
}

func GetReadMeConfigStr(name string) string {
	for _, module := range GetConfig().ReadMe.Modules {
		nameKey := util.GetValueFromBracket(module)
		if len(nameKey) < 2 {
			continue
		}
		if nameKey[0] == name {
			return nameKey[1]
		}
	}
	return ""
}

func GetConfig() compatibility.Config {
	return getConfigOnce()
}

func UpdateConfig() {
	compatilityList := []compatibility.ConfigCompatibility{
		compatibility.NewTomlConfig(fmt.Sprintf("%s/.unemeta_cli.yaml", util.HomeDir()), fmt.Sprintf("%s/.unemeta_cli.toml", util.HomeDir())),
		compatibility.NewDbDocConfigCompatility(defaultConfig()),
		compatibility.NewConfigUpdate(fmt.Sprintf("%s/.unemeta_cli.toml", util.HomeDir()), toml.Marshal),
	}
	for _, compatility := range compatilityList {
		compatility.Compatility(&config)
	}
}

func UpdateCompletion(cmd *cobra.Command) {
	fPathStr := util.MustCmdZshWithOutput("echo $fpath")
	fPathList := strings.Split(strings.TrimSpace(fPathStr), " ")
	fmt.Printf("Please select one path below to update zsh completion func.\n")
	fmt.Printf("[0]Skip\n")
	for i, p := range fPathList {
		fmt.Printf("[%d]%s\n", i+1, p)
	}
	var num int
	for {
		_, err := fmt.Scanln(&num)
		if err != nil || num < 0 || num > len(fPathList) {
			if err != nil && err.Error() == "unexpected newline" {
				continue
			}
			fmt.Printf("Please insert a number range [0, %d].\n", len(fPathList))
			continue
		}
		break
	}
	if num != 0 {
		fPath := fPathList[num-1]
		completionFilePath := path.Join(fPath, "_une")
		util.MustCmd(fmt.Sprintf("sudo touch %s && sudo chmod 777 %s", completionFilePath, completionFilePath))
		file, err := os.OpenFile(completionFilePath, os.O_WRONLY|os.O_TRUNC, 0777)
		if err != nil {
			log.Printf("open file %s err:%+v", completionFilePath, err)
			return
		}
		err = cmd.Root().GenZshCompletion(file)
		if err != nil {
			log.Printf("gen zsh completion err:%+v\n", err)
			return
		}
		fmt.Printf("gen zsh completion success, maybe you should restart shell to take effect.\n")
	} else {
		return
	}
}
