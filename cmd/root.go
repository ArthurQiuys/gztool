package cmd

import (
	"git.unemeta.com/Backstage/une/cmd/handle_orient"
	"git.unemeta.com/Backstage/une/cmd/object_orient"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "une",
	Short: "unemeta cli helper",
	Long:  `general used cli wrapper for unemeta project`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(handle_orient.GetCmds()...)
	rootCmd.AddCommand(object_orient.GetCmds()...)
}
