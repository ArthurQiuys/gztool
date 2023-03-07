package object_orient

import (
	"fmt"
	"path"

	"git.unemeta.com/Backstage/une/src/app"
	"git.unemeta.com/Backstage/une/src/util"
	"github.com/spf13/cobra"
)

func getAppCmds() []*cobra.Command {
	return []*cobra.Command{rpcCmd(), sqlNewCmd(), sqlUpCmd(), sqlModelCmd(), sqlDownCmd(), runCmd(), docCmd(), yApiCmd(), readMeCmd()}
}

func GetCmds() []*cobra.Command {
	var cmd []*cobra.Command
	for _, a := range app.GetAllApp() {
		useStr := func() string {
			if len(a.Name()) > 3 {
				return a.Name()[:3]
			}
			return a.Name()
		}()
		appCmd := &cobra.Command{
			Use:   useStr,
			Short: fmt.Sprintf("%s module handles wrapper", a.Name()),
		}
		appCmd.AddCommand(getAppCmds()...)
		if util.AssetExist(path.Join(a.Path(), "api")) {
			appCmd.AddCommand(apiCmd("api"))
		}
		if util.AssetExist(path.Join(a.Path(), "mq")) {
			appCmd.AddCommand(apiCmd("mq"))
		}
		if util.AssetExist(path.Join(a.Path(), "job")) {
			appCmd.AddCommand(apiCmd("job"))
		}
		if util.AssetExist(path.Join(a.Path(), "scheduler")) {
			appCmd.AddCommand(apiCmd("scheduler"))
		}
		cmd = append(cmd, appCmd)
	}
	return cmd
}
