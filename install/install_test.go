package install

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gregex"
	"github.com/onism68/helmOnap/utils"
	"github.com/onism68/helmOnap/vars"
	"testing"
)

func TestRunInstall(t *testing.T) {
	vars.MasterIp = "127.0.0.1"
	vars.NodeIps = []string{"-l", "onism", "-L", "192.168.2.249"}
	RunInstall()
}

func Test_runInNode(t *testing.T) {
	vars.SSHConfig.User = "root"
	vars.SSHConfig.Password = "0000"
	vars.MasterIp = "172.21.80.101"
	sshMaster := utils.SSH{
		User:     "root",
		Password: "0000",
		//PkFile:     "/root/.ssh/id_rsa",
		PkPassword: "",
		Timeout:    nil,
	}
	//sshMaster.CmdInServer(vars.MasterIp, "helm serve &")
	//return
	nodes := []string{"172.21.80.101", "172.21.80.102"}
	nodes = append(nodes, vars.MasterIp)
	lenNodes := len(nodes)
	fmt.Println(nodes)
	fmt.Println(nodes[:lenNodes-1])
	fmt.Println(nodes[lenNodes:])
	nodes = append(nodes[:lenNodes-1], nodes[lenNodes:]...)
	fmt.Println(nodes)
	return
	//list := getNodesSource(nodes, vars.WorkSpace+"docker/", "tar")
	//DockerLoader(nodes, list)
	//return
	//runInNode([]string{"172.21.80.101"}, CdCom(vars.WorkSpace))

	t1 := sshMaster.CmdInServer(vars.MasterIp, "cd workSpaceTmp/docker && ls")
	//fmt.Println(string(t1))
	matchString, err2 := gregex.MatchAll(`([\S]+\.tar)`, t1)
	if err2 != nil {
		panic(err2)
	}
	//fmt.Println(matchString)
	var fileList []string
	for _, item := range matchString {
		//glog.Info(index)
		fileList = append(fileList, string(item[0]))
	}
	fmt.Println(fileList)
	return
	//runInNode([]string{"172.21.80.101"}, WgetCom("172.21.80.1:8199/main.exe"))

	// 检查tiller服务是否就绪
	tmp := sshMaster.CmdInServer(vars.MasterIp, "kubectl get pods -n kube-system | grep tiller")
	fmt.Println(tmp)
	readyList, err := gregex.MatchString(`(\d+)\/(\d+)`, string(tmp))
	if err != nil {
		panic(err)
	}
	readyList[1] = "1"
	readyList[2] = "2"
	for readyList[1] != readyList[2] {
		glog.Info("tiller 服务未就绪")
		tillerInfo := sshMaster.CmdInServer(vars.MasterIp, "kubectl get pods -n kube-system | grep tiller")
		readyList, err = gregex.MatchString(`(\d+)\/(\d+)`, string(tillerInfo))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("已就绪")
	fmt.Println(readyList)
}
