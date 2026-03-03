package version

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBuildInfo_ReturnsDefaultValues(t *testing.T) {
	info := GetBuildInfo()

	assert.Equal(t, "v0.0.0-unset", info.Version)
	assert.Equal(t, "unset", info.Commit)
}

func TestGetBuildInfo_PlatformMatchesRuntime(t *testing.T) {
	info := GetBuildInfo()

	expected := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	assert.Equal(t, expected, info.Platform)
}

func TestBuildInfo_String_ContainsAllFields(t *testing.T) {
	info := GetBuildInfo()
	s := info.String()

	assert.Contains(t, s, info.Version)
	assert.Contains(t, s, info.Commit)
	assert.Contains(t, s, info.Platform)
}

func TestBuildInfo_String_Format(t *testing.T) {
	info := GetBuildInfo()
	s := info.String()

	// Expected format: "v0.0.0-unset (linux/amd64) [unset]"
	assert.True(t, strings.HasPrefix(s, info.Version), "should start with version")
	assert.Contains(t, s, "("+info.Platform+")")
	assert.Contains(t, s, "["+info.Commit+"]")
}
