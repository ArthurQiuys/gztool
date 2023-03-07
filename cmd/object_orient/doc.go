package object_orient

import (
	"git.unemeta.com/Backstage/une/src/app"
	"git.unemeta.com/Backstage/une/src/config"
	"github.com/spf13/cobra"
)

func docCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "doc",
		Short: "only generate swagger json.",
		Run: func(cmd *cobra.Command, args []string) {
			appName := cmd.Parent().Name()
			app.GetApp(appName).Doc(config.DocNone)
		},
	}
}

func yApiCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "yapi",
		Short: "yapi import wrapper",
		Run: func(cmd *cobra.Command, args []string) {
			appName := cmd.Parent().Name()
			app.GetApp(appName).Doc(config.DocYApi)
		},
	}
}

func readMeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "readme",
		Short: "readme import wrapper",
		Run: func(cmd *cobra.Command, args []string) {
			appName := cmd.Parent().Name()
			app.GetApp(appName).Doc(config.DocReadMe)
		},
	}
}
