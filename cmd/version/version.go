package version

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wallanaq/hew/internal/version"
)

type options struct {
	short bool
	json  bool
}

// NewVersionCommand creates and returns the version command for the CLI application.
func NewVersionCommand() *cobra.Command {
	opts := &options{}

	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "Print the version number of Hew",
		Aliases: []string{"v"},
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return printVersion(opts)
		},
	}

	f := versionCmd.Flags()
	f.BoolVar(&opts.short, "short", false, "Print just the version number.")
	f.BoolVar(&opts.json, "json", false, "Print the version in JSON format.")

	return versionCmd
}

// printVersion handles the logic of printing the version information
// based on the provided flags (e.g., --short, --json).
func printVersion(opts *options) error {
	info := version.GetBuildInfo()

	switch {
	case opts.json:
		return json.NewEncoder(os.Stdout).Encode(info)
	case opts.short:
		fmt.Println(info.Version)
	default:
		fmt.Printf("hew version %s\n", info)
	}

	return nil
}
