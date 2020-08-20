package utils

import (
	"bufio"
	"fmt"
	"github.com/gogf/gf/os/glog"
	"os/exec"
)

func RunCmd(name string, args []string) {
	cmdr := exec.Command(name, args...)
	pipe, _:= cmdr.StdoutPipe()
	if err := cmdr.Start();  err != nil {
		glog.Error("some error : %s", err.Error())
	}
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Println(line)
		line, err = reader.ReadString('\n')
	}
}
