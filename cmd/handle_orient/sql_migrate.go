package handle_orient

import (
	"fmt"

	"git.unemeta.com/Backstage/une/src/config"

	"git.unemeta.com/Backstage/une/src/app"
	"github.com/spf13/cobra"
)

var migrateEnv string
var openInDataGrip bool
var useCache bool

const (
	test = "test"
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql-migrate wrapper.",
}

func AppModuleValidFunc(_ *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
	if len(args) < 1 {
		return app.GetNames(), cobra.ShellCompDirectiveNoFileComp
	}
	return []string{}, cobra.ShellCompDirectiveNoFileComp
}

func getEnv() config.SqlEnvEnum {
	if migrateEnv == test {
		return config.DbTest
	}
	return config.DbLocal
}

var sqlNewCmd = &cobra.Command{
	Use:               "new",
	Short:             "sql-migrate new wrapper.",
	Long:              `First args should be app module name, second should be migrate file name.`,
	Args:              cobra.MatchAll(cobra.ExactArgs(2)),
	ValidArgsFunction: AppModuleValidFunc,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			return
		}
		app := app.GetApp(args[0])
		if app != nil {
			app.MigrateNew(args[1], openInDataGrip)
		} else {
			fmt.Println("app module not correct:", args[0])
		}
	},
}

var sqlUpCmd = &cobra.Command{
	Use:               "up",
	Short:             "sql-migrate up wrapper.",
	Long:              `Only one arg should be app module name, [-e/local]=[test/local] [-c/true]=[false/true].`,
	Args:              cobra.MatchAll(cobra.ExactArgs(1)),
	ValidArgsFunction: AppModuleValidFunc,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			return
		}
		app := app.GetApp(args[0])
		if app != nil {
			app.MigrateUp(getEnv())
		} else {
			fmt.Println("app module not correct:", args[0])
		}
	},
}

var sqlDownCmd = &cobra.Command{
	Use:               "down",
	Short:             "sql-migrate down wrapper.",
	Long:              `Only one arg should be app module name, have one -e env flag.`,
	Args:              cobra.MatchAll(cobra.ExactArgs(1)),
	ValidArgsFunction: AppModuleValidFunc,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			return
		}
		app := app.GetApp(args[0])
		if app != nil {
			app.MigrateDown(getEnv())
		} else {
			fmt.Println("app module not correct:", args[0])
		}
	},
}

var sqlModelCmd = &cobra.Command{
	Use:               "mo",
	Short:             "goctl model migrate.",
	Long:              `Only one arg should be app module name, [-e/local]=[test/local] [-c/true]=[false/true].`,
	Args:              cobra.MatchAll(cobra.ExactArgs(1)),
	ValidArgsFunction: AppModuleValidFunc,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			return
		}
		app := app.GetApp(args[0])
		if app != nil {
			app.MigrateModel(getEnv(), useCache)
		} else {
			fmt.Println("app module not correct:", args[0])
		}
	},
}

func getSqlMigrateCmd() *cobra.Command {
	sqlUpCmd.Flags().StringVarP(&migrateEnv, "env", "e", "local", "should be local or test")

	sqlModelCmd.Flags().StringVarP(&migrateEnv, "env", "e", "local", "should be local or test")
	sqlModelCmd.Flags().BoolVarP(&useCache, "cache", "c", true, "use cache generate model, default true")

	sqlDownCmd.Flags().StringVarP(&migrateEnv, "env", "e", "local", "should be local or test")
	sqlNewCmd.Flags().BoolVarP(&openInDataGrip, "open", "o", false, "use datagrip open migrate file.")

	sqlCmd.AddCommand(sqlNewCmd, sqlUpCmd, sqlDownCmd)
	return sqlCmd
}
