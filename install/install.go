package install

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/onism68/helmOnap/utils"
)

func RunInstall() {
	port := g.Cfg().GetString("port")

	nodes := utils.ParseIPs(NodeIps)
	pkgUrl := fmt.Sprintf("http://%s:%s/%s", MasterIp, port, PkgName)
	runInNode(nodes, pkgUrl)
}

func runInNode(nodes []string, pkgUrl string) {
	for _, nodeIp := range nodes {
		args := []string{nodeIp}
		extraName := []string{"ping"}
		//extraArgs := [][]string{[]string{pkgUrl}}
		extraArgs := [][]string{[]string{"-c", "10", "192.168.2.1"}, []string{"pwd"}}
		utils.RunCmd("ssh", args, extraName, extraArgs)
	}

}
