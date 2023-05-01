package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	_ "unsafe"
)

type Version struct {
	Ver       string
	Revision  string
	Timestamp string

	Compiler  string
	BuildTime string
}

//go:linkname ver main.__version__
var ver Version

func init() {
	ver.BuildTime = time.Now().Format(time.DateTime)

	if ver.Timestamp == "" {
		ver.Timestamp = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(ver.BuildTime, "-", ""), ":", ""), " ", "_")
	}

	if ver.Revision == "" {
		ver.Revision = fmt.Sprintf("%s", ver.Timestamp)
	}

	if ver.Ver == "" {
		ver.Ver = fmt.Sprintf("v0.0.0-%s", ver.Revision)
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		ver.Compiler = runtime.Version()
		return
	}

	ver.Compiler = info.GoVersion
}

func Short() string {
	return ver.Ver
}

func Full() Version {
	return ver
}
