package version

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	internalversion "github.com/wallanaq/hew/internal/version"
)

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
		cmd.SetArgs([]string{})
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
		cmd.SetArgs([]string{"--short"})
		err := cmd.Execute()
		assert.NoError(t, err)
	})

	assert.Equal(t, info.Version, strings.TrimSpace(output))
}

func TestNewVersionCommand_ExecuteJsonFlag(t *testing.T) {
	output := captureStdout(t, func() {
		cmd := NewVersionCommand()
		cmd.SetErr(new(bytes.Buffer))
		cmd.SetArgs([]string{"--json"})
		err := cmd.Execute()
		assert.NoError(t, err)
	})

	var info internalversion.BuildInfo
	err := json.Unmarshal([]byte(output), &info)
	assert.NoError(t, err)
	assert.NotEmpty(t, info.Version)
}
