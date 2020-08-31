package install

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gregex"
	"github.com/onism68/helmOnap/utils"
	"github.com/onism68/helmOnap/vars"
	"os"
	"time"
)

func RunInstall() {

	// 创建工作目录
	//utils.CmdsInLocal("mkdir", []string{vars.WorkSpace}, nil, nil)
	//// master解压包
	//runInMaster("touch", []string{"test.test"})
	//runInMaster("tar", []string{"-zxvf", vars.PkgName, "-C", vars.WorkSpace})
	//// 安装helm
	//runInMaster("cp", []string{"./helm/helm", "/usr/local/bin/"})
	//// 配置权限
	//runInMaster("chmod", []string{"+x", "/usr/local/bin/helm"})
	//
	//runInMaster("kubectl", []string{"create", "serviceaccount --namespace=kube-system tiller"})
	//runInMaster("kubectl", []string{"create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller"})
	////helm init
	//runInMaster("helm", []string{"init", "--upgrade -i registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.16.6 --stable-repo-url http://mirror.azure.cn/kubernetes/charts/ --service-account=tiller"})
	//runInMaster("helm", []string{"repo remove stable"})

	//runInMaster("cp", []string{"-r", "./helm/plugins", "~/.helm/"})
	//s := utils.CmdToString("helm", fmt.Sprintf("serve %s", "&"))
	//glog.Println(s)
	//runInMaster("helm", []string{fmt.Sprintf("serve %s", "&")})

	//return
	// 配置本地文件服务器地址
	port := g.Cfg().GetString("port")
	nodes := utils.ParseIPs(vars.NodeIps)
	pkgUrl := fmt.Sprintf("http://%s:%s/%s", vars.MasterIp, port, vars.PkgName)
	//通过ssh链接本地
	sshMaster := utils.SSH{
		User:     vars.SSHConfig.User,
		Password: vars.SSHConfig.Password,
		//PkFile:     "/root/.ssh/id_rsa",
		PkPassword: "",
		Timeout:    nil,
	}
	//创建工作目录
	sshMaster.CmdInMaster(Mkdir(vars.WorkSpace))

	sshMaster.CmdInMaster(TarX(vars.PkgPath+vars.PkgName, vars.WorkSpace))
	// oom make好的文件
	sshMaster.CmdInMaster(Mkdir(vars.WorkSpace + "oom"))
	sshMaster.CmdInMaster(TarX(vars.WorkSpace+"oom-onap-f-IfNot.tar.gz", vars.WorkSpace+"oom"))
	// 安装helm，配置权限
	sshMaster.CmdInMaster(Cpr(vars.WorkSpace+"/helm/helm", "/usr/local/bin/"))
	sshMaster.CmdInMaster(Chmod("+x", "/usr/local/bin/helm"))

	//资源分发，下面安装tiller需要镜像，所以先进行资源分发
	// 下载包
	// 创建目录
	runInNode(nodes, Mkdir(vars.WorkSpace))
	// 下载包
	runInNode(nodes, Wget(pkgUrl, vars.PkgName))
	// 解压包
	runInNode(nodes, TarX(vars.PkgName, vars.WorkSpace))
	// 先load tiller镜像
	loadTiller := "docker load -i " + vars.WorkSpace + "/helm/tiller-v2.16.6.tar"
	// 安装配置tiller相关
	sshMaster.CmdInMaster(loadTiller)
	runInNode(nodes, loadTiller)
	sshMaster.CmdInMaster("kubectl create serviceaccount --namespace=kube-system tiller")
	sshMaster.CmdInMaster("kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin " +
		"--serviceaccount=kube-system:tiller")
	sshMaster.CmdInMaster("helm init --upgrade -i registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.16.6 " +
		"--stable-repo-url http://mirror.azure.cn/kubernetes/charts/ --service-account=tiller")

	// 检查tiller服务是否就绪
	tillerInfo := sshMaster.CmdInServer(vars.MasterIp, "kubectl get pods -n kube-system | grep tiller")
	readyList, err := gregex.MatchString(`(\d+)\/(\d+)`, string(tillerInfo))
	if err != nil {
		panic(err)
	}
	forTimes := 0
	for readyList[1] != readyList[2] {
		glog.Info(utils.CmdTips("tiller 服务未就绪"))
		tillerInfo := sshMaster.CmdInServer(vars.MasterIp, "kubectl get pods -n kube-system | grep tiller")
		readyList, err = gregex.MatchString(`(\d+)\/(\d+)`, string(tillerInfo))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
		forTimes++
		if forTimes > 60 {
			glog.Error(utils.CmdTips("tiller 服务持续未就绪，exit"))
			os.Exit(vars.ErrorExitOSCase)
		}
	}
	glog.Info(utils.CmdTips("tiller 服务已就绪，进行后续操作"))
	//sshMaster.CmdInMaster("helm repo remove stable")
	sshMaster.CmdInMaster(Cpr(vars.WorkSpace+"helm/plugins", "~/.helm/"))
	//sshMaster.CmdInMaster("helm create namespace onap")
	// 创建namespace
	sshMaster.CmdInMaster("kubectl create namespace onap")

	//sshMaster.CmdInMaster(fmt.Sprintf("docker load -i %s", vars.WorkSpace+"docker/docker.tar || true"))
	//sshMaster.CmdInMaster(Chmod("+x", vars.WorkSpace + "docker/docker.sh"))
	//sshMaster.CmdInMaster(fmt.Sprintf("/bin/sh %s", vars.WorkSpace+"docker/docker.sh"))
	// todo 暂时无法后台运行，重新考虑
	//sshMaster.CmdInServer(vars.MasterIp, "helm serve >> out.txt 2>&1 &")

	////sshMaster.CmdInMaster(HelmInstall(VFC + " --set global.masterPassword=onap"))
	//return

	// 获取某目录下所需要的文件
	nodes = append(nodes, vars.MasterIp)
	list := getNodesSource(nodes, vars.WorkSpace+"docker/", "tar")
	DockerLoader(nodes, list)
	nodes = append(nodes[:len(nodes)-1])
	glog.Info(utils.CmdTips("install msb"))
	sshMaster.CmdInMaster(InstallMsb())

	glog.Info(utils.CmdTips("install cassandra"))
	sshMaster.CmdInMaster(InstallCassandra())

	glog.Info(utils.CmdTips("install AAI"))
	sshMaster.CmdInMaster(InstallAAI())

	glog.Info(utils.CmdTips("install VFC"))
	sshMaster.CmdInMaster(InstallVFC())

	glog.Info(utils.CmdTips("install Modeling"))
	sshMaster.CmdInMaster(InstallModeling())

	glog.Info(utils.CmdTips("install Multicloud"))
	sshMaster.CmdInMaster(InstallMulticloud())

	glog.Info(utils.CmdTips("install ESR"))
	sshMaster.CmdInMaster(InstallEsr())

	glog.Info(utils.CmdTips("install UUI"))
	sshMaster.CmdInMaster(InstallUUI())

}

func CleanWorkSpace() {
	nodes := utils.ParseIPs(vars.NodeIps)
	nodes = append(nodes, vars.MasterIp)
	runInNode(nodes, "rm -rf %s"+vars.WorkSpace)
}

func runInMaster(name string, args []string) {
	utils.CmdInLocal(name, args)
}

func runInNode(nodes []string, arg string) {
	for _, nodeIp := range nodes {
		ssh := utils.SSH{
			User:       vars.SSHConfig.User,
			Password:   vars.SSHConfig.Password,
			PkFile:     vars.SSHConfig.PrivateKeyPath,
			PkPassword: "",
			Timeout:    nil,
		}
		//_ = ssh.CmdInServer(nodeIp, Mkdir(vars.WorkSpace))
		err := ssh.CmdAsync(nodeIp, arg)
		if err != nil {
			glog.Error(err.Error())
		}
	}
}

/**
cd 目录
suffix 后缀
dataList 节点列表所获取到的文件列表[][]
*/
func getNodesSource(nodes []string, cd string, suffix string) (dataList [][]string) {
	for _, nodeIp := range nodes {
		ssh := utils.SSH{
			User:       vars.SSHConfig.User,
			Password:   vars.SSHConfig.Password,
			PkFile:     vars.SSHConfig.PrivateKeyPath,
			PkPassword: "",
			Timeout:    nil,
		}
		//_ = ssh.CmdInServer(nodeIp, Mkdir(vars.WorkSpace))
		data := ssh.CmdInServer(nodeIp, fmt.Sprintf("cd %s && ls", cd))
		list := utils.FindFileList(data, suffix)
		dataList = append(dataList, list)
	}
	return dataList
}

/**

 */

func DockerLoader(nodes []string, tarList [][]string) {
	for index, nodeIp := range nodes {
		ssh := utils.SSH{
			User:       vars.SSHConfig.User,
			Password:   vars.SSHConfig.Password,
			PkFile:     vars.SSHConfig.PrivateKeyPath,
			PkPassword: "",
			Timeout:    nil,
		}
		go func(nodeIp string, index int) {
			for _, tar := range tarList[index] {
				err := ssh.CmdAsync(nodeIp, fmt.Sprintf("docker load -i %s", vars.WorkSpace+"/docker/"+tar))
				if err != nil {
					glog.Error(err)
					os.Exit(vars.ErrorExitOSCase)
				}
			}
		}(nodeIp, index)

	}
}
