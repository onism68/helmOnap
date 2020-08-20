package install

var (
	MasterIp string
	NodeIps []string
	PkgPath = "/root/helmOnap/"
	PkgName = ""
	ErrorExitOSCase = -1 // 错误直接退出类型

)

var SSHConfig struct{
	User string
	Password string
	PrivateKeyPath string
}
