package object_orient

import (
	"fmt"

	"git.unemeta.com/Backstage/une/src/app"
	"git.unemeta.com/Backstage/une/src/config"
	"github.com/spf13/cobra"
)

var migrateEnv string
var openInDataGrip bool
var useCache bool

const (
	test = "test"
)

func getEnv() config.SqlEnvEnum {
	if migrateEnv == test {
		return config.DbTest
	}
	return config.DbLocal
}

func sqlNewCmd() *cobra.Command {
	sqlCmd := &cobra.Command{
		Use:   "new",
		Short: "sql-migrate new wrapper.",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				return
			}
			appModule := app.GetApp(cmd.Parent().Name())
			if appModule != nil {
				appModule.MigrateNew(args[0], openInDataGrip)
			}
		},
	}
	sqlCmd.Flags().BoolVarP(&openInDataGrip, "open", "o", false, "use datagrip open migrate file")
	return sqlCmd
}

func sqlUpCmd() *cobra.Command {
	sqlCmd := &cobra.Command{
		Use:   "up",
		Short: "sql-migrate up wrapper.",
		Run: func(cmd *cobra.Command, args []string) {
			appModule := app.GetApp(cmd.Parent().Name())
			if appModule != nil {
				appModule.MigrateUp(getEnv())
			}
		},
	}
	sqlCmd.Flags().StringVarP(&migrateEnv, "env", "e", "local", "sql environment, default local")
	return sqlCmd
}

func sqlDownCmd() *cobra.Command {
	sqlCmd := &cobra.Command{
		Use:   "down",
		Short: "sql-migrate down wrapper.",
		Run: func(cmd *cobra.Command, args []string) {
			appModule := app.GetApp(cmd.Parent().Name())
			if appModule != nil {
				appModule.MigrateDown(getEnv())
			}
		},
	}
	sqlCmd.Flags().StringVarP(&migrateEnv, "env", "e", "local", "sql environment, default local")
	return sqlCmd
}

func sqlModelCmd() *cobra.Command {
	sqlCmd := &cobra.Command{
		Use:   "mo",
		Short: "goctl model migrate.",
		Long:  `[-e/local]=[test/local] [-c/true]=[false/true].`,
		Run: func(cmd *cobra.Command, args []string) {
			getApp := app.GetApp(cmd.Parent().Name())
			if getApp != nil {
				getApp.MigrateModel(getEnv(), useCache)
			} else {
				fmt.Println("app module not correct:", cmd.Parent().Name())
			}
		},
	}
	sqlCmd.Flags().BoolVarP(&useCache, "cache", "c", true, "use cache generate model, default true")
	sqlCmd.Flags().StringVarP(&migrateEnv, "env", "e", "local", "sql environment, default local")
	return sqlCmd
}
