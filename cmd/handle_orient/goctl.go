package handle_orient

import (
	"sync"

	"git.unemeta.com/Backstage/une/cmd/valid"

	"git.unemeta.com/Backstage/une/src/app"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:               "api",
	Short:             "goctl api wrapper",
	ValidArgsFunction: valid.AppModulesValidFunc,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		for _, arg := range args {
			app := app.GetApp(arg)
			if app != nil {
				wg.Add(1)
				go func() {
					app.Api("api")
					wg.Done()
				}()
			}
		}
		wg.Wait()
	},
}

var rpcCmd = &cobra.Command{
	Use:               "rpc",
	Short:             "goctl rpc wrapper",
	ValidArgsFunction: valid.AppModulesValidFunc,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		for _, arg := range args {
			app := app.GetApp(arg)
			if app != nil {
				wg.Add(1)
				go func() {
					app.Rpc()
					wg.Done()
				}()
			}
		}
		wg.Wait()
	},
}
