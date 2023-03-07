package handle_orient

import (
	"fmt"

	"git.unemeta.com/Backstage/une/src/app"
	"git.unemeta.com/Backstage/une/src/config"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "go run wrapper",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) <= 0 {
			return []string{"api", "rpc", "mq"}, cobra.ShellCompDirectiveNoFileComp
		}
		if len(args) == 1 {
			return app.GetNames(), cobra.ShellCompDirectiveNoFileComp
		}
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			return
		}
		app := app.GetApp(args[1])
		if app != nil {
			app.Run(args[0])
		} else {
			panic(fmt.Sprintf("app module:%s not found in project dir:%s", args[1], config.GetConfig().ProjectDir))
		}
	},
}
