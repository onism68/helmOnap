package install

import (
	"fmt"
	"github.com/onism68/helmOnap/vars"
)

var (
	mkdir = "mkdir -p %s"
	chmod = "chmod %s %s"
	cd    = "cd %s"
	wget  = "wget %s -O %s"
	cp    = "cp -r %s %s"
	tarx  = "tar -zxvf %s -C %s"
)

type HelmInstallFile string

const VFC HelmInstallFile = "vfc"
const AAI HelmInstallFile = "aai"
const MSB HelmInstallFile = "msb"
const MODELING HelmInstallFile = "modeling"
const CASSANDRA HelmInstallFile = "cassandra"
const UUI HelmInstallFile = "uui"
const MULTICLOUD HelmInstallFile = "multicloud"
const MARIADB HelmInstallFile = "mariadb-galera"
const ESR HelmInstallFile = "esr"

func Mkdir(dir string) string {
	return fmt.Sprintf(mkdir, dir)
}

func Cd(path string) string {
	return fmt.Sprintf(cd, path)
}

func Wget(url string, name string) string {
	return fmt.Sprintf(wget, url, name)
}

func TarX(tarPath string, toPath string) string {
	return fmt.Sprintf(tarx, tarPath, toPath)
}

func Cpr(filePath string, toPath string) string {
	return fmt.Sprintf(cp, filePath, toPath)
}

func Chmod(mod string, path string) string {
	return fmt.Sprintf(chmod, mod, path)
}

func HelmInstall(install HelmInstallFile, arg string) string {
	// todo 自定义版本
	return fmt.Sprintf("helm --namespace onap install %soom/%s-6.0.0.tgz %s", vars.WorkSpace, install, arg)
}

func InstallVFC() string {
	return HelmInstall(VFC, "--name dev-vfc")
}

func InstallModeling() string {
	return HelmInstall(MODELING, "--name dev-modeling --set global.masterPassword=onap")
}

func InstallAAI() string {
	return HelmInstall(AAI, "--name dev-aai")
}

func InstallMsb() string {
	return HelmInstall(MSB, "--name dev-msb")
}

func InstallCassandra() string {
	return HelmInstall(CASSANDRA, "--name dev-cassandra")
}

func InstallMulticloud() string {
	return HelmInstall(MULTICLOUD, "--name dev-multicloud")
}

func InstallEsr() string {
	return HelmInstall(ESR, "--name dev-esr")
}
func InstallUUI() string {
	return HelmInstall(UUI, "--name dev-uui")
}
