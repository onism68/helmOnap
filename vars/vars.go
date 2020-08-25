package vars

import "time"

var (
	MasterIp        string
	NodeIps         []string
	PkgPath         = "/root/workSpaceTmp/"
	PkgName         = ""
	WorkSpace       = "/tmp/workSpaceTmp/"
	ErrorExitOSCase = -1 // 错误直接退出类型

)

var SSHConfig struct {
	User           string
	Password       string
	PrivateKeyPath string
}

type SSH struct {
	User       string
	Password   string
	PkFile     string
	PkPassword string
	Timeout    *time.Duration
}
