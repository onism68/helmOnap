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
const CATALOG HelmInstallFile = "modeling"
const CASSANDRA HelmInstallFile = "cassandra"
const UUI HelmInstallFile = "uui"
const MULTICLOUD HelmInstallFile = "multicloud"
const mariadb HelmInstallFile = "mariadb-galera"
const ESR HelmInstallFile = "esr"

func MkdirCom(dir string) string {
	return fmt.Sprintf(mkdir, dir)
}

func CdCom(path string) string {
	return fmt.Sprintf(cd, path)
}

func WgetCom(url string, name string) string {
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

func HelmInstall(install HelmInstallFile) string {
	return fmt.Sprintf("helm --namespace onap install %s/onap/charts/%s", vars.WorkSpace, install)
}