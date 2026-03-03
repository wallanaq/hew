package version

import (
	"fmt"
	"runtime"
)

var (
	version = "v0.0.0-unset"
	commit  = "unset"
)

// BuildInfo holds all the build-time information.
type BuildInfo struct {
	Version  string
	Commit   string
	Platform string
}

// Get returns the build information of the application.
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version:  version,
		Commit:   commit,
		Platform: fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// String returns a formatted string representation of the BuildInfo.
func (i BuildInfo) String() string {
	return fmt.Sprintf("%s (%s) [%s]", i.Version, i.Platform, i.Commit)
}
