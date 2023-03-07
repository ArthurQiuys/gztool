package object_orient

import (
	"fmt"

	"git.unemeta.com/Backstage/une/src/app"
	"github.com/spf13/cobra"
)

func apiCmd(n string) *cobra.Command {
	return &cobra.Command{
		Use:   n,
		Short: fmt.Sprintf("generate %s api by goctl api.", n),
		Run: func(cmd *cobra.Command, args []string) {
			appModule := app.GetApp(cmd.Parent().Name())
			if appModule != nil {
				appModule.Api(n)
			} else {
				fmt.Println("app module not found.")
			}
		},
	}
}

func rpcCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "rpc",
		Short: "goctl rpc wrapper",
		Run: func(cmd *cobra.Command, args []string) {
			appModule := app.GetApp(cmd.Parent().Name())
			if appModule != nil {
				appModule.Rpc()
			} else {
				fmt.Println("app module not found.")
			}
		},
	}
}
