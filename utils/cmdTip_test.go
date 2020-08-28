package utils

import (
	"fmt"
	"testing"
)

func TestCmdTips(t *testing.T) {
	fmt.Println(CmdTips("testtesttest"))
	fmt.Println(CmdTips("1111111111"))
	fmt.Println(len("                              "))
}
