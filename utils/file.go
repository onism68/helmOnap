package utils

import (
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gregex"
	"github.com/onism68/helmOnap/vars"
)

func FindFileList(byte []byte, suffix string) []string {
	match, err := gregex.MatchAll(fmt.Sprintf(`([\S]+\.%s)`, suffix), byte)
	if err != nil {
		glog.Error(err.Error())
	}
	var fileList []string
	for _, item := range match {
		fileList = append(fileList, string(item[1]))
	}
	return fileList
}

// 以下两个功能缺陷
//todo 重写
func ReadFile2List(file string) error {
	return gfile.ReadByteLines(file, byte2List)
}

func byte2List(byte []byte) {
	vars.ImagesList = append(vars.ImagesList, string(byte))
}
