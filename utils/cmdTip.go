package utils

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
)

func CmdTips(tip string) string {
	l := "------------------------------ \n"
	t := "                               "

	lenTip := showStrLen(tip)
	glog.Println(lenTip)
	if lenTip <= 28 {
		glog.Info(15 - lenTip/2)
		tip = string(t[1:15-lenTip/2]) + tip
	}
	return fmt.Sprintf("\n"+l+"%s\n"+l, tip)
}

func showStrLen(s string) int {
	sl := 0
	rs := []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			sl++
		} else {
			sl += 2
		}
	}
	return sl
}
