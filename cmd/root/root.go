package root

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/wallanaq/hew/cmd/version"
	internalversion "github.com/wallanaq/hew/internal/version"
)

type options struct {
	debug bool
}

// NewRootComannd creates and returns the root command for the CLI application.
func NewRootCommand() *cobra.Command {
	opts := &options{}

	rootCmd := &cobra.Command{
		Use:           "hew",
		Short:         "Sharp tools for developers. Built for the terminal.",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       internalversion.GetBuildInfo().String(),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if opts.debug {
				slog.SetLogLoggerLevel(slog.LevelDebug)
				slog.Debug("debug mode enabled")
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	f := rootCmd.PersistentFlags()
	f.BoolVar(&opts.debug, "debug", false, "enable debug mode")

	rootCmd.AddCommand(
		version.NewVersionCommand(),
	)

	return rootCmd
}
