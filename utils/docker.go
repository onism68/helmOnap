package utils

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/grand"
	"github.com/onism68/helmOnap/vars"
	"strings"
)

func PullOrSaveImage() {
	sshMaster := SSH{
		User:     vars.SSHConfig.User,
		Password: vars.SSHConfig.Password,
		//PkFile:     "/root/.ssh/id_rsa",
		PkPassword: "",
		Timeout:    nil,
	}
	glog.Info(sshMaster)
	err := ReadFile2List(vars.ImagesListFile, &vars.ImagesList)
	if err != nil {
		glog.Error(err.Error())
	}
	lens := len(vars.ImagesList)
	ifMkdir := true
	var tmpList []string
	for index, item := range vars.ImagesList {
		glog.Info(vars.DockerPull)
		if vars.DockerPull {
			glog.Infof("----pulled %d, sum %d----", index, lens)
			sshMaster.CmdInMaster(fmt.Sprintf("docker pull %s", item))
			tmpList = append(tmpList, item+"\n")
		} else {
			if ifMkdir {
				sshMaster.CmdInMaster("mkdir images")
				ifMkdir = false
			}
			var nameTmp string
			if strings.Contains(item, "/") {
				nameTmpList := strings.Split(item, "/")
				nameTmp = nameTmpList[len(nameTmpList)-1]
				nameTmp = strings.Replace(nameTmp, ":", "-", -1)
			} else {
				nameTmp = item
				nameTmp = strings.Replace(nameTmp, ":", "-", -1)
			}
			glog.Infof("-----image ifno: ( %s )", nameTmp)
			saveImageName := nameTmp + grand.Letters(5)
			sshMaster.CmdInMaster(fmt.Sprintf("docker save %s > ./images/%s.tar", item, saveImageName))
			tmpList = append(tmpList, saveImageName+"\n")
		}
	}
	glog.Info(tmpList)
}
