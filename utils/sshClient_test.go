package utils

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"testing"
)

func TestSSH_Connect(t *testing.T) {
	ssh := SSH{
		User:     "root",
		Password: "0222",
		//PkFile:     "/root/.ssh/id_rsa",
		PkPassword: "",
		Timeout:    nil,
	}
	conn, err := ssh.Connect("172.21.80.102")
	if err != nil {
		glog.Error(err.Error())
	}
	output, err := conn.Output("ping -c 10 172.21.80.1")
	fmt.Println(string(output))
}
