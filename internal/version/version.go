package version

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"golang.org/x/mod/semver"
)

var (
	version = "v0.0.0-unset"
	commit  = "unset"
)

var latestReleaseURL = "https://api.github.com/repos/wallanaq/hew/releases/latest"

// BuildInfo holds all the build-time information.
type BuildInfo struct {
	Version  string `json:"version"`
	Commit   string `json:"commit"`
	Platform string `json:"platform"`
}

// Get returns the build information of the application.
func GetBuildInfo() *BuildInfo {
	return &BuildInfo{
		Version:  version,
		Commit:   commit,
		Platform: fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// String returns a formatted string representation of the BuildInfo.
func (i BuildInfo) String() string {
	return fmt.Sprintf("%s (%s) [%s]", i.Version, i.Platform, i.Commit)
}

// UpdateInfo holds the result of an update check.
type UpdateInfo struct {
	CurrentVersion string
	LatestVersion  string
}

// HasUpdate reports whether a newer version is available than the current one.
func (i *UpdateInfo) HasUpdate() bool {
	return semver.Compare(i.LatestVersion, i.CurrentVersion) > 0
}

// ReleaseInfo holds the Github release info.
type ReleaseInfo struct {
	TagName     string `json:"tag_name"`
	CreatedAt   string `json:"created_at"`
	PublishedAt string `json:"published_at"`
}

// CheckForUpdates queries the GitHub Releases API for the latest release.
// The provided context controls the timeout and cancellation.
func CheckForUpdates(ctx context.Context) (*UpdateInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, latestReleaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	// #nosec G704 -- URL is a hardcoded constant, not user input
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var release ReleaseInfo
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &UpdateInfo{
		CurrentVersion: version,
		LatestVersion:  release.TagName,
	}, nil
}
