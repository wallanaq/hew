package root

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRootCommand_ReturnsValidCommand(t *testing.T) {
	cmd := NewRootCommand()
	assert.NotNil(t, cmd)
}

func TestNewRootCommand_Use(t *testing.T) {
	cmd := NewRootCommand()
	assert.Equal(t, "hew", cmd.Use)
}

func TestNewRootCommand_SilencesUsageAndErrors(t *testing.T) {
	cmd := NewRootCommand()
	assert.True(t, cmd.SilenceUsage)
	assert.True(t, cmd.SilenceErrors)
}

func TestNewRootCommand_HasVersionSet(t *testing.T) {
	cmd := NewRootCommand()
	assert.NotEmpty(t, cmd.Version)
}

func TestNewRootCommand_RegistersVersionSubcommand(t *testing.T) {
	cmd := NewRootCommand()

	var found bool
	for _, sub := range cmd.Commands() {
		if sub.Use == "version" {
			found = true
			break
		}
	}

	assert.True(t, found, "subcommand 'version' should be registered")
}

func TestNewRootCommand_ExecuteNoArgs(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.NoError(t, err)
}
