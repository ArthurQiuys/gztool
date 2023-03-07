package handle_orient

import (
	"github.com/spf13/cobra"
)

func GetCmds() []*cobra.Command {
	return []*cobra.Command{getSqlMigrateCmd(), apiCmd, rpcCmd, runCmd}
}
