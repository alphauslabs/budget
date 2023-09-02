package source

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/alphauslabs/blue-internal-go/budget/v1"
	"github.com/alphauslabs/bluectl/pkg/logger"
	"github.com/alphauslabs/budget/pkg/connection"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <id> [yyyymm]",
		Short: "List the latest query source for alerts",
		Long: `List the latest query source for alerts.

The <id> is required and should be an account id.
If [yyyymm] is not set, it defaults to the current month.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			con, err := connection.New(ctx)
			if err != nil {
				logger.Errorf("connection.New failed: %v", err)
				return
			}

			if len(args) == 0 {
				logger.Errorf("<id> is required.")
				return
			}

			id := args[0]
			month := time.Now().UTC().Format("200601")
			if len(args) > 1 {
				month = args[1]
			}

			defer con.Close()
			client := budget.NewBudgetClient(con)
			req := budget.ListAlertSourceRequest{
				Vendor: "aws",
				AwsOptions: &budget.ListAlertSourceRequest_AwsOptions{
					Account: id,
					Month:   month,
				},
			}

			stream, err := client.ListAlertSource(ctx, &req)
			if err != nil {
				logger.Errorf("ListAlertSource failed: %v", err)
				return
			}

		loop:
			for {
				v, err := stream.Recv()
				switch {
				case err == io.EOF:
					break loop
				case err != nil && err != io.EOF:
					logger.Error(err)
					break loop
				}

				b, _ := json.Marshal(v)
				logger.Info(string(b))
			}
		},
	}

	cmd.Flags().SortFlags = false
	return cmd
}
