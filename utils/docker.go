package utils

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/onism68/helmOnap/vars"
)

func PullImage() {
	sshMaster := SSH{
		User:     vars.SSHConfig.User,
		Password: vars.SSHConfig.Password,
		//PkFile:     "/root/.ssh/id_rsa",
		PkPassword: "",
		Timeout:    nil,
	}
	glog.Info(sshMaster)
	err := ReadFile2List(vars.ImagesListFile)
	if err != nil {
		glog.Error(err.Error())
	}
	lens := len(vars.ImagesList)
	for index, item := range vars.ImagesList {
		glog.Infof("----pulled %d, sum %d----", index, lens)
		sshMaster.CmdInMaster(fmt.Sprintf("docker pull %s", item))
	}
}

func SaveImage() {

}