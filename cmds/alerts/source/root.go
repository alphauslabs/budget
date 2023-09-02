package source

import (
	"github.com/alphauslabs/bluectl/pkg/logger"
	"github.com/spf13/cobra"
)

func SourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source",
		Short: "Subcommand for budget alert source operations",
		Long:  `Subcommand for budget alert source operations.`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("see -h for more information")
		},
	}

	cmd.Flags().SortFlags = false
	cmd.AddCommand(ListCmd())
	return cmd
}
