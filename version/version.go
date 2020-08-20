package version

import (
"fmt"
	"github.com/gogf/gf/os/gtime"
	"runtime"
)

var (
	Version   = "0.0.1"
	Build     = ""
	BuildTime = gtime.Datetime()
	VersionStr = fmt.Sprintf(" version: %v \n build: %v %v \n Build Time: %v", Version, Build, runtime.Version(), BuildTime)
)


