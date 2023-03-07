package valid

import (
	"git.unemeta.com/Backstage/une/src/app"
	"git.unemeta.com/Backstage/une/src/util"
	"github.com/spf13/cobra"
	"strings"
)

func AppModulesValidFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var res []string
	for _, s := range app.GetNames() {
		var already = false
		for _, arg := range args {
			if strings.HasPrefix(s, arg) {
				already = true
			}
		}
		if !already {
			res = append(res, s)
		}
	}
	return res, cobra.ShellCompDirectiveNoFileComp
}

func RunModuleValidFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) < 1 {
		appName := cmd.Parent().Name()
		res := util.GetSubDirectories(app.GetApp(appName).Path())
		return res, cobra.ShellCompDirectiveNoFileComp
	}
	return []string{}, cobra.ShellCompDirectiveNoFileComp
}
