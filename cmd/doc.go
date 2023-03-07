package cmd

import (
	"sync"

	"git.unemeta.com/Backstage/une/src/app"
	"git.unemeta.com/Backstage/une/src/config"
	"github.com/spf13/cobra"
)

var yApiCmd = &cobra.Command{
	Use:   "yapi",
	Short: "yapi import wrapper",
	Run: func(cmd *cobra.Command, args []string) {
		apps := app.GetAllApp()
		var wg sync.WaitGroup
		for _, appModule := range apps {
			wg.Add(1)
			go func(a app.App) {
				defer wg.Done()
				a.Doc(config.DocYApi)
			}(appModule)
		}
		wg.Wait()
	},
}

var readMeCmd = &cobra.Command{
	Use:   "readme",
	Short: "readme import wrapper",
	Run: func(cmd *cobra.Command, args []string) {
		apps := app.GetAllApp()
		var wg sync.WaitGroup
		for _, appModule := range apps {
			wg.Add(1)
			go func(a app.App) {
				defer wg.Done()
				a.Doc(config.DocReadMe)
			}(appModule)
		}
		wg.Wait()
	},
}

var noDocCmd = &cobra.Command{
	Use:   "no",
	Short: "open generate swagger json.",
	Run: func(cmd *cobra.Command, args []string) {
		apps := app.GetAllApp()
		var wg sync.WaitGroup
		for _, appModule := range apps {
			wg.Add(1)
			go func(a app.App) {
				defer wg.Done()
				a.Doc(config.DocNone)
			}(appModule)
		}
		wg.Wait()
	},
}
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "goctl doc goctl-swagger wrapper.",
}

func init() {
	docCmd.AddCommand(yApiCmd, readMeCmd, noDocCmd)
	rootCmd.AddCommand(docCmd)
}
