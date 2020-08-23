package utils

import (
	"bufio"
	"bytes"
	"github.com/gogf/gf/os/glog"
	"github.com/onism68/helmOnap/vars"
	"io"
	"os"
	"os/exec"
)

func CmdInLocal(name string, args []string) {
	CmdsInLocal(name, args, nil, nil)
}

func CmdsInLocal(name string, args []string, extraNames []string, extraArgs [][]string) {
	cmdr := exec.Command(name, args...)
	cmdr.Dir = vars.WorkSpace
	glog.Printf("name %s, args %s", name, args)
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
			CmdsInLocal(extra, extraArgs[index], nil, nil)
		}
	}
}

//CmdToString is exec on os , return result
func CmdToString(name string, arg ...string) string {
	glog.Infof("[os]exec cmd is : %s, %s", name, arg)
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	err := cmd.Run()
	if err != nil {
		glog.Errorf("[os]os call error. %s", err)
		return ""
	}
	return b.String()
}

//Cmd is in host exec cmd
func (ss *SSH) CmdInServer(host string, cmd string) []byte {
	glog.Infof("[ssh][%s] %s", host, cmd)
	session, err := ss.Connect(host)
	defer func() {
		if r := recover(); r != nil {
			glog.Errorf("[ssh][%s]Error create ssh session failed,%s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	defer session.Close()
	b, err := session.CombinedOutput(cmd)
	glog.Debugf("[ssh][%s]command result is: %s", host, string(b))
	defer func() {
		if r := recover(); r != nil {
			glog.Errorf("[ssh][%s]Error exec command failed: %s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	return b
}

func readPipe(host string, pipe io.Reader, isErr bool) {
	for {
		r := bufio.NewReader(pipe)
		line, _, err := r.ReadLine()
		if line == nil {
			return
		} else if err != nil {
			glog.Infof("[%s] %s", host, line)
			glog.Errorf("[ssh] [%s] %s", host, err)
			return
		} else {
			if isErr {
				glog.Errorf("[%s] %s", host, line)
			} else {
				glog.Infof("[%s] %s", host, line)
			}
		}
	}
}

func (ss *SSH) CmdAsync(host string, cmd string) error {
	glog.Infof("[ssh][%s] %s", host, cmd)
	session, err := ss.Connect(host)
	if err != nil {
		glog.Errorf("[ssh][%s]Error create ssh session failed,%s", host, err)
		return err
	}
	defer session.Close()
	stdout, err := session.StdoutPipe()
	if err != nil {
		glog.Errorf("[ssh][%s]Unable to request StdoutPipe(): %s", host, err)
		return err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		glog.Errorf("[ssh][%s]Unable to request StderrPipe(): %s", host, err)
		return err
	}
	if err := session.Start(cmd); err != nil {
		glog.Errorf("[ssh][%s]Unable to execute command: %s", host, err)
		return err
	}
	doneout := make(chan bool, 1)
	doneerr := make(chan bool, 1)
	go func() {
		readPipe(host, stderr, true)
		doneerr <- true
	}()
	go func() {
		readPipe(host, stdout, false)
		doneout <- true
	}()
	<-doneerr
	<-doneout
	return nil
}

func (ss *SSH) CmdInMaster(args string) {
	err := ss.CmdAsync(vars.MasterIp, args)
	if err != nil {
		glog.Error(err)
	}
}
