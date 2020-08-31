package utils

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"testing"
)

func TestCmdTips(t *testing.T) {
	glog.Info(CmdTips("install msb"))
	fmt.Println(CmdTips("testtesttest"))
	fmt.Println(CmdTips("1111111111"))
	fmt.Println(CmdTips("测试测试啊啊啊啊啊"))
	fmt.Println(len("                              "))
}
