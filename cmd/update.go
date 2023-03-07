package cmd

import (
	"fmt"
	"sync"
	"time"

	"git.unemeta.com/Backstage/une/src/config"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update une cli config file, goctl-swagger, goctl_template",
	Run: func(cmd *cobra.Command, args []string) {
		// update config.
		config.UpdateConfig()
		// update completion
		config.UpdateCompletion(cmd)
		fmt.Printf("while update cache dependencies after a second.\n")
		time.Sleep(time.Second)
		// update dependencies
		{
			var wg sync.WaitGroup
			// update goctl-swagger
			wg.Add(1)
			go func() {
				defer wg.Done()
				UpdateGoCtlSwagger()
			}()
			wg.Add(1)
			// update goctl-template
			go func() {
				defer wg.Done()
				UpdateGoCtlTemplate()
			}()
			wg.Wait()
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
