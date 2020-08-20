package install

import (
	"bytes"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/onism68/helmOnap/utils"
	"os/exec"
)

func RunInstall() {
	port := g.Cfg().GetString("port")
	nodes := utils.ParseIPs(NodeIps)
	pkgUrl := fmt.Sprintf("http://%s:%s/%s", MasterIp, port, PkgName)
}

func runInNode(nodes []string) {
	var outInfo bytes.Buffer
	for _, nodeIp := range nodes {
		args := []string{nodeIp}
		utils.RunCmd("ssh", args)
	}

}