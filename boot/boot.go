package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func Boot() {
	glog.Info("配置信息")
	c := g.Cfg()
	s := g.Server()
	// glog配置
	logpath := c.GetString("setting.logpath")
	glog.SetPath(logpath)
	glog.SetStdoutPrint(true)

	// 静态文件配置
	s.SetIndexFolder(true)
	s.SetServerRoot(c.GetString("Pkg.PkgPath"))
	//s.AddSearchPath("/Users/john/Documents")
	s.SetPort(c.GetInt("port"))
	// 开启日志
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	glog.Info("配置完成，准备启动")
	s.Run()
}
