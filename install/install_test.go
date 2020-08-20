package install

import "testing"

func TestRunInstall(t *testing.T) {
	MasterIp = "127.0.0.1"
	NodeIps = []string{"-l", "onism", "-L", "192.168.2.249"}
	RunInstall()
}
