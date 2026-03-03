package version

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/wallanaq/hew/internal/doc"
	"github.com/wallanaq/hew/internal/version"
)

type options struct {
	json          bool
	short         bool
	noUpdateCheck bool
}

// NewVersionCommand creates and returns the version command for the CLI application.
func NewVersionCommand() *cobra.Command {
	opts := &options{}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hew",
		Long: doc.Description(`
			Print the current version of the Hew CLI, including build metadata.

			By default, the output includes the version string and commit information.
			Use --short to print only the version number, or --json to get structured output.

			An update check is performed automatically unless --no-update-check is passed.`),
		Example: doc.Example(`
			# Print the full version information
			hew version

			# Print only the version number
			hew version --short

			# Print version in JSON format
			hew version --json

			# Skip the update check
			hew version --no-update-check`),
		Args:    cobra.NoArgs,
		Aliases: []string{"v"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), opts)
		},
	}

	f := versionCmd.Flags()
	f.BoolVar(&opts.short, "short", false, "Print just the version number.")
	f.BoolVar(&opts.json, "json", false, "Print the version in JSON format.")
	f.BoolVar(&opts.noUpdateCheck, "no-update-check", false, "Disable update check.")

	return versionCmd
}

// run prints the version and, unless disabled by opts, checks for newer releases.
// Update check errors are silenced so they never break the command.
func run(ctx context.Context, opts *options) error {
	if err := printVersion(opts); err != nil {
		return fmt.Errorf("print version: %w", err)
	}

	if opts.noUpdateCheck {
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	info, err := version.CheckForUpdates(ctx)
	if err != nil {
		slog.Debug("failed to check updates", slog.Any("error", err))
		return nil
	}

	if info.HasUpdate() {
		fmt.Fprintf(os.Stderr, "A newer version is available: %s -> %s\n", info.CurrentVersion, info.LatestVersion)
	}

	return nil
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
