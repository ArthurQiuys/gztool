package object_orient

import (
	"git.unemeta.com/Backstage/une/cmd/valid"
	"git.unemeta.com/Backstage/une/src/app"
	"github.com/spf13/cobra"
)

func runCmd() *cobra.Command {
	return &cobra.Command{
		Use:               "run",
		Short:             "go run wrapper",
		Args:              cobra.MatchAll(cobra.ExactArgs(1)),
		ValidArgsFunction: valid.RunModuleValidFunc,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				return
			}
			appModule := app.GetApp(cmd.Parent().Name())
			if appModule != nil {
				appModule.Run(args[0])
			}
		},
	}
}
