package utils

import (
	"testing"
)

func TestRunCmd(t *testing.T) {
	CmdsInLocal("echo", []string{"111"},
		[]string{"echo", "echo", "echo"},
		[][]string{[]string{"222"}, []string{"333"}, []string{"444"}},
	)
}

func TestSSH_CmdAsync(t *testing.T) {
	ssh := SSH{
		User:       "root",
		Password:   "0222",
		PkFile:     "/root/.ssh/id_rsa",
		PkPassword: "",
		Timeout:    nil,
	}
	ssh.CmdAsync("172.21.80.101", "ping -c 10 172.21.80.1")
}
