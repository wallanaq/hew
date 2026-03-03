package version

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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

func TestUpdateInfo_HasUpdate_WhenNewerVersion(t *testing.T) {
	info := &UpdateInfo{CurrentVersion: "v1.0.0", LatestVersion: "v2.0.0"}
	assert.True(t, info.HasUpdate())
}

func TestUpdateInfo_HasUpdate_WhenSameVersion(t *testing.T) {
	info := &UpdateInfo{CurrentVersion: "v1.0.0", LatestVersion: "v1.0.0"}
	assert.False(t, info.HasUpdate())
}

func TestUpdateInfo_HasUpdate_WhenOlderVersion(t *testing.T) {
	info := &UpdateInfo{CurrentVersion: "v2.0.0", LatestVersion: "v1.0.0"}
	assert.False(t, info.HasUpdate())
}

func TestCheckForUpdates_ReturnsNewerVersion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ReleaseInfo{TagName: "v99.0.0"})
	}))
	defer srv.Close()

	original := latestReleaseURL
	latestReleaseURL = srv.URL
	t.Cleanup(func() { latestReleaseURL = original })

	info, err := CheckForUpdates(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "v99.0.0", info.LatestVersion)
	assert.True(t, info.HasUpdate())
}

func TestCheckForUpdates_WhenUpToDate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ReleaseInfo{TagName: version})
	}))
	defer srv.Close()

	original := latestReleaseURL
	latestReleaseURL = srv.URL
	t.Cleanup(func() { latestReleaseURL = original })

	info, err := CheckForUpdates(context.Background())
	assert.NoError(t, err)
	assert.False(t, info.HasUpdate())
}

func TestCheckForUpdates_WhenNonOKStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	original := latestReleaseURL
	latestReleaseURL = srv.URL
	t.Cleanup(func() { latestReleaseURL = original })

	_, err := CheckForUpdates(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected status")
}

func TestCheckForUpdates_WhenInvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "not json")
	}))
	defer srv.Close()

	original := latestReleaseURL
	latestReleaseURL = srv.URL
	t.Cleanup(func() { latestReleaseURL = original })

	_, err := CheckForUpdates(context.Background())
	assert.Error(t, err)
}

func TestCheckForUpdates_WhenContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := CheckForUpdates(ctx)
	assert.Error(t, err)
}
