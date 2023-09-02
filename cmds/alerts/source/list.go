package source

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/alphauslabs/blue-internal-go/budget/v1"
	"github.com/alphauslabs/bluectl/pkg/logger"
	"github.com/alphauslabs/budget/pkg/connection"
	"github.com/olekukonko/tablewriter"
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

			render := true
			table := tablewriter.NewWriter(os.Stdout)
			table.SetBorder(false)
			table.SetAutoWrapText(false)
			table.SetHeaderLine(false)
			table.SetColumnSeparator("")
			table.SetTablePadding("  ")
			table.SetNoWhiteSpace(true)
			table.SetColumnAlignment([]int{
				tablewriter.ALIGN_DEFAULT,
				tablewriter.ALIGN_DEFAULT,
				tablewriter.ALIGN_RIGHT,
				tablewriter.ALIGN_RIGHT,
				tablewriter.ALIGN_RIGHT,
				tablewriter.ALIGN_RIGHT,
			})

			table.Append([]string{
				"ACCT",
				"DATE",
				"TRUEUNBLENDED",
				"UNBLENDED",
				"TU_DIFF",
				"U_DIFF",
			})

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

				if v.Aws == nil {
					continue
				}

				add := []string{
					id,
					v.Aws.Date,
					fmt.Sprintf("%f", v.Aws.TrueUnblended),
					fmt.Sprintf("%f", v.Aws.Unblended),
				}

				if v.Aws.TrueUnblendedDiff > 0.001 {
					add = append(add, fmt.Sprintf("%f", v.Aws.TrueUnblendedDiff))
				} else {
					add = append(add, "-")
				}

				if v.Aws.UnblendedDiff > 0.001 {
					add = append(add, fmt.Sprintf("%f", v.Aws.UnblendedDiff))
				} else {
					add = append(add, "-")
				}

				table.Append(add)
			}

			if render {
				table.Render()
			}
		},
	}

	cmd.Flags().SortFlags = false
	return cmd
}
