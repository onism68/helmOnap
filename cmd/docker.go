package cmd

import (
	"github.com/onism68/helmOnap/utils"
	"github.com/onism68/helmOnap/vars"
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "docker",
	Long:  "docker pull",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		utils.PullOrSaveImage(vars.DockerPull)
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.Flags().StringVar(&vars.ImagesListFile, "file", "./images.list", "pull image list")
	dockerCmd.Flags().StringVar(&vars.SSHConfig.User, "user", "root", "servers user name for ssh")
	dockerCmd.Flags().StringVar(&vars.SSHConfig.Password, "password", "0000", "servers user password for ssh")
	dockerCmd.Flags().StringVar(&vars.MasterIp, "masterIp", "192.168.0.2", "kubernetes multi-masters ex. 192.168.0.2")
	dockerCmd.Flags().BoolVar(&vars.DockerPull, "pull", true, "true is pull, false is save, default true")

	dockerCmd.MarkFlagRequired("file")
	dockerCmd.MarkFlagRequired("user")
	dockerCmd.MarkFlagRequired("masterIp")
	dockerCmd.MarkFlagRequired("password")
}
