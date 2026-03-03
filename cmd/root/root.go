package root

import "github.com/spf13/cobra"

// NewRootComannd creates and returns the root command for the CLI application.
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "hew",
		Short:         "Sharp tools for developers. Hew through repetitive tasks.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	return rootCmd
}
