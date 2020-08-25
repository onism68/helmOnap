package utils

import (
	"github.com/gogf/gf/os/glog"
	"testing"
)

func TestReadFile(t *testing.T) {
	file := "../images.list"
	err := ReadFile2List(file)
	if err != nil {
		glog.Error(err)
	}
}
