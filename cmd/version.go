package cmd

import (
	"fmt"
	"github.com/onism68/helmOnap/version"
	"github.com/spf13/cobra"
)


var versionCmd = &cobra.Command{
	Use: "version",
	Short: "version",
	Long: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.VersionStr)
	},
}


func init() {
	rootCmd.AddCommand(versionCmd)
}