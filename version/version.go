package version

import (
	"fmt"
	"runtime"
)

var (
	Version    = "0.0.1"
	Build      = "linux_amd64"
	BuildTime  = "2020-08-28 19:00"
	VersionStr = fmt.Sprintf(" version: %v \n build: %v %v \n Build Time: %v", Version, Build, runtime.Version(), BuildTime)
)
