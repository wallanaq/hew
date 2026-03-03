package version

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	internalversion "github.com/wallanaq/hew/internal/version"
)

// mockTransport intercepts all HTTP requests and returns a configurable response.
type mockTransport struct {
	tagName    string
	statusCode int
}

func (m *mockTransport) RoundTrip(*http.Request) (*http.Response, error) {
	body, _ := json.Marshal(map[string]string{"tag_name": m.tagName})
	return &http.Response{
		StatusCode: m.statusCode,
		Body:       io.NopCloser(strings.NewReader(string(body))),
		Header:     make(http.Header),
	}, nil
}

func withMockUpdateCheck(t *testing.T, tagName string) {
	t.Helper()
	original := http.DefaultTransport
	http.DefaultTransport = &mockTransport{tagName: tagName, statusCode: http.StatusOK}
	t.Cleanup(func() { http.DefaultTransport = original })
}

func withFailingUpdateCheck(t *testing.T) {
	t.Helper()
	original := http.DefaultTransport
	http.DefaultTransport = &mockTransport{statusCode: http.StatusInternalServerError}
	t.Cleanup(func() { http.DefaultTransport = original })
}

func captureStderr(t *testing.T, f func()) string {
	t.Helper()

	old := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stderr = w
	f()
	w.Close()
	os.Stderr = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func captureStdout(t *testing.T, f func()) string {
	t.Helper()

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestNewVersionCommand_ReturnsValidCommand(t *testing.T) {
	cmd := NewVersionCommand()
	assert.NotNil(t, cmd)
}

func TestNewVersionCommand_Use(t *testing.T) {
	cmd := NewVersionCommand()
	assert.Equal(t, "version", cmd.Use)
}

func TestNewVersionCommand_HasAliasV(t *testing.T) {
	cmd := NewVersionCommand()
	assert.Contains(t, cmd.Aliases, "v")
}

func TestNewVersionCommand_HasShortFlag(t *testing.T) {
	cmd := NewVersionCommand()
	flag := cmd.Flags().Lookup("short")
	assert.NotNil(t, flag, "flag --short should be registered")
	assert.Equal(t, "false", flag.DefValue)
}

func TestNewVersionCommand_HasJsonFlag(t *testing.T) {
	cmd := NewVersionCommand()
	flag := cmd.Flags().Lookup("json")
	assert.NotNil(t, flag, "flag --json should be registered")
	assert.Equal(t, "false", flag.DefValue)
}

func TestPrintVersion_Default(t *testing.T) {
	info := internalversion.GetBuildInfo()

	output := captureStdout(t, func() {
		err := printVersion(&options{})
		assert.NoError(t, err)
	})

	assert.Contains(t, output, "hew version")
	assert.Contains(t, output, info.String())
}

func TestPrintVersion_Short(t *testing.T) {
	info := internalversion.GetBuildInfo()

	output := captureStdout(t, func() {
		err := printVersion(&options{short: true})
		assert.NoError(t, err)
	})

	assert.Equal(t, info.Version, strings.TrimSpace(output))
}

func TestPrintVersion_JSON(t *testing.T) {
	output := captureStdout(t, func() {
		err := printVersion(&options{json: true})
		assert.NoError(t, err)
	})

	var info internalversion.BuildInfo
	err := json.Unmarshal([]byte(output), &info)
	assert.NoError(t, err)
	assert.NotEmpty(t, info.Version)
	assert.NotEmpty(t, info.Commit)
	assert.NotEmpty(t, info.Platform)
}

func TestNewVersionCommand_ExecuteDefault(t *testing.T) {
	output := captureStdout(t, func() {
		cmd := NewVersionCommand()
		cmd.SetErr(new(bytes.Buffer))
		cmd.SetArgs([]string{"--no-update-check"})
		err := cmd.Execute()
		assert.NoError(t, err)
	})

	assert.Contains(t, output, "hew version")
}

func TestNewVersionCommand_ExecuteShortFlag(t *testing.T) {
	info := internalversion.GetBuildInfo()

	output := captureStdout(t, func() {
		cmd := NewVersionCommand()
		cmd.SetErr(new(bytes.Buffer))
		cmd.SetArgs([]string{"--short", "--no-update-check"})
		err := cmd.Execute()
		assert.NoError(t, err)
	})

	assert.Equal(t, info.Version, strings.TrimSpace(output))
}

func TestNewVersionCommand_ExecuteJsonFlag(t *testing.T) {
	output := captureStdout(t, func() {
		cmd := NewVersionCommand()
		cmd.SetErr(new(bytes.Buffer))
		cmd.SetArgs([]string{"--json", "--no-update-check"})
		err := cmd.Execute()
		assert.NoError(t, err)
	})

	var info internalversion.BuildInfo
	err := json.Unmarshal([]byte(output), &info)
	assert.NoError(t, err)
	assert.NotEmpty(t, info.Version)
}

func TestNewVersionCommand_HasNoUpdateCheckFlag(t *testing.T) {
	cmd := NewVersionCommand()
	flag := cmd.Flags().Lookup("no-update-check")
	assert.NotNil(t, flag, "flag --no-update-check should be registered")
	assert.Equal(t, "false", flag.DefValue)
}

func TestRun_WithNoUpdateCheck_SkipsUpdateCheck(t *testing.T) {
	var runErr error
	output := captureStdout(t, func() {
		runErr = run(context.Background(), &options{noUpdateCheck: true})
	})
	assert.NoError(t, runErr)
	assert.Contains(t, output, "hew version")
}

func TestRun_PrintsUpdateNoticeToStderr(t *testing.T) {
	withMockUpdateCheck(t, "v99.0.0")

	var runErr error
	stderr := captureStderr(t, func() {
		captureStdout(t, func() {
			runErr = run(context.Background(), &options{})
		})
	})
	assert.NoError(t, runErr)
	assert.Contains(t, stderr, "v99.0.0")
}

func TestRun_SilentWhenUpToDate(t *testing.T) {
	withMockUpdateCheck(t, "v0.0.0-unset")

	var runErr error
	stderr := captureStderr(t, func() {
		captureStdout(t, func() {
			runErr = run(context.Background(), &options{})
		})
	})
	assert.NoError(t, runErr)
	assert.Empty(t, stderr)
}

func TestRun_SilentOnUpdateCheckError(t *testing.T) {
	withFailingUpdateCheck(t)

	var runErr error
	stderr := captureStderr(t, func() {
		captureStdout(t, func() {
			runErr = run(context.Background(), &options{})
		})
	})
	assert.NoError(t, runErr)
	assert.Empty(t, stderr)
}
