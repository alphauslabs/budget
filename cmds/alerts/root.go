package alerts

import (
	"github.com/alphauslabs/bluectl/pkg/logger"
	"github.com/alphauslabs/budget/cmds/alerts/source"
	"github.com/spf13/cobra"
)

func AlertsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alerts",
		Short: "Subcommand for budget alert operations",
		Long:  `Subcommand for budget alert operations.`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("see -h for more information")
		},
	}

	cmd.Flags().SortFlags = false
	cmd.AddCommand(source.SourceCmd())
	return cmd
}
