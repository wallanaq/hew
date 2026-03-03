package root

import (
	"github.com/spf13/cobra"
	"github.com/wallanaq/hew/cmd/version"
	internalversion "github.com/wallanaq/hew/internal/version"
)

// NewRootComannd creates and returns the root command for the CLI application.
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "hew",
		Short:         "Sharp tools for developers. Hew through repetitive tasks.",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       internalversion.GetBuildInfo().String(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.AddCommand(
		version.NewVersionCommand(),
	)

	return rootCmd
}
