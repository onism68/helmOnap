package utils

import "testing"

func TestRunCmd(t *testing.T) {
	RunCmd("echo", []string{"111"},
		[]string{"echo", "echo", "echo"},
		[][]string{[]string{"222"}, []string{"333"}, []string{"444"}},
	)
}
