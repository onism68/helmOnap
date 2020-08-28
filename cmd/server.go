package cmd

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/onism68/helmOnap/boot"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long:  "run a file server in current dir",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		if err := g.Cfg().Set("Pkg.PkgPath", "./"); err == nil {
			boot.Boot()
		} else {
			glog.Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
