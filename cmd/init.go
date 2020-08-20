package cmd

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/onism68/helmOnap/install"
	"github.com/spf13/cobra"
	"path"
)

var initCmd = &cobra.Command{
	Use: "init",
	Short: "init",
	Long: "init",
	Run: func(cmd *cobra.Command, args []string) {
	//
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		glog.Info("写入临时路径")
		install.PkgPath, install.PkgName = path.Split(install.PkgPath)
		err := g.Cfg().Set("Pkg.PkgPath", install.PkgPath)
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
	initCmd.Flags().StringVar(&install.SSHConfig.User, "user", "root", "servers user name for ssh")
	initCmd.Flags().StringVar(&install.SSHConfig.Password, "password", "", "servers user password for ssh")
	initCmd.Flags().StringVar(&install.SSHConfig.PrivateKeyPath, "PrivateKeyPath", "/root/.ssh/id_rsa", "private key for ssh")
	
	initCmd.Flags().StringVar(&install.PkgPath, "pkgPath", "/root/helmOnap/xxx.tar.gz", "path of pkg ex. /root/helmOnap/test.tar.gz")
	initCmd.MarkFlagRequired("pkgPath")
	initCmd.Flags().StringVar(&install.MasterIp, "MasterIp", "192.168.0.2", "kubernetes multi-masters ex. 192.168.0.2")
	initCmd.Flags().StringSliceVar(&install.NodeIps, "NodeIps", []string{}, "kubernetes multi-nodes ex. 192.168.0.5, 192.168.0.6")


}