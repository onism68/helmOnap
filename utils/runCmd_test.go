package utils

import (
	"github.com/gogf/gf/os/glog"
	"github.com/onism68/helmOnap/vars"
	"testing"
)

func TestRunCmd(t *testing.T) {
	CmdsInLocal("echo", []string{"111"},
		[]string{"echo", "echo", "echo"},
		[][]string{[]string{"222"}, []string{"333"}, []string{"444"}},
	)
}

func TestSSH_CmdAsync(t *testing.T) {
	ssh := SSH{
		User:       "root",
		Password:   "0000",
		PkFile:     "/root/.ssh/id_rsa",
		PkPassword: "",
		Timeout:    nil,
	}
	vars.MasterIp = "172.21.80.101"
	ssh.CmdAsync("172.21.80.101", "ping -c 3 172.21.80.1")
	glog.Infof("\x1b[%d;%dm%s\x1b[0m", 44, 30, "hello")
	//ssh.CmdInMaster(fmt.Sprintf("docker load -i %s", vars.WorkSpace+"docker/docker.tar || true"))
	//ssh.CmdInMaster(fmt.Sprintf("sh %s", vars.WorkSpace+"docker/docker.sh"))

}
