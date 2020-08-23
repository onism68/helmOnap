package cmd

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/onism68/helmOnap/install"
	"github.com/onism68/helmOnap/vars"
	"github.com/spf13/cobra"
	"path"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init",
	Long:  "init",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		glog.Info("写入临时路径")
		vars.PkgPath, vars.PkgName = path.Split(vars.PkgPath)
		err := g.Cfg().Set("Pkg.PkgPath", vars.PkgPath)
		if err != nil {
			glog.Error("写入临时路径出错：%s", err.Error())
		}
		// 后台起一个server，用于包分发
		install.RunServer()

		//	开始安装
		install.RunInstall()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// todo 读取配置文件，更新
	initCmd.Flags().StringVar(&vars.SSHConfig.User, "user", "root", "servers user name for ssh")
	initCmd.Flags().StringVar(&vars.SSHConfig.Password, "password", "0222", "servers user password for ssh")
	initCmd.Flags().StringVar(&vars.SSHConfig.PrivateKeyPath, "PrivateKeyPath", "/root/.ssh/id_rsa", "private key for ssh")

	initCmd.Flags().StringVar(&vars.PkgPath, "pkgPath", "/root/helmOnap/xxx.tar.gz", "path of pkg ex. /root/helmOnap/test.tar.gz")
	initCmd.Flags().StringVar(&vars.MasterIp, "masterIp", "192.168.0.2", "kubernetes multi-masters ex. 192.168.0.2")
	initCmd.Flags().StringSliceVar(&vars.NodeIps, "nodeIps", []string{}, "kubernetes multi-nodes ex. 192.168.0.5, 192.168.0.6")

	initCmd.MarkFlagRequired("masterIp")
	initCmd.MarkFlagRequired("nodeIps")
	initCmd.MarkFlagRequired("pkgPath")

}
