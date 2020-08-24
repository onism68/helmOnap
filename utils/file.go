package utils

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gregex"
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
