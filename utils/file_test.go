package utils

import (
	"github.com/gogf/gf/os/glog"
	"github.com/onism68/helmOnap/vars"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	file := "../images.list"
	err := ReadFile2List(file)
	if err != nil {
		glog.Error(err)
	}
	i := 0
	for _, item := range vars.ImagesList {
		nameAndVer := strings.Split(item, ":")
		glog.Info(nameAndVer)
		nameTmpList := strings.Split(nameAndVer[0], "/")
		name := nameTmpList[len(nameTmpList)-1]
		glog.Info(name)
		if i > 10 {
			return
		}
		i++
	}

}
