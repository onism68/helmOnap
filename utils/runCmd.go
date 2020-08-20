package utils

import (
	"bufio"
	"github.com/gogf/gf/os/glog"
	"io"
	"log"
	"os"
	"os/exec"
)

func RunCmd(name string, args []string, extraNames []string, extraArgs [][]string) {
	cmdr := exec.Command(name, args...)
	glog.Info(name, args)
	pipeReader, _ := cmdr.StdoutPipe()
	if err := cmdr.Start(); err != nil {
		glog.Errorf("some error : %s", err.Error())
	}
	defer pipeReader.Close()
	reader := bufio.NewReader(pipeReader)
	line, err := reader.ReadString('\n')
	for err == nil {
		glog.Info(line)
		line, err = reader.ReadString('\n')
		for index, extra := range extraNames {
			RunCmd(extra, extraArgs[index], nil, nil)
		}
	}
}

func RunCmd2() {
	ps := exec.Command("ps", "aux")
	grep := exec.Command("grep", "go")
	reader, err := ps.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()
	grep.Stdin = reader
	writer, err := grep.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}
	defer writer.Close()
	ps.Start()
	defer ps.Wait()
	grep.Start()
	defer grep.Wait()
	io.Copy(os.Stdout, writer)
}
