package blob

import (
	"github.com/carolynvs/az-cli/cmd/az/command"
	"github.com/carolynvs/az-cli/cmd/az/storage/blob/lease"
	"github.com/spf13/cobra"
)

type SubCommand struct {
	*command.Context
	*cobra.Command
}

func NewSubCommand(cxt *command.Context) *SubCommand {
	c := &SubCommand{
		Context: cxt,
		Command: &cobra.Command{
			Use:   "blob",
			Short: "blob commands",
		},
	}

	c.AddCommand(
		lease.NewSubCommand(cxt).Command,
		NewUploadBatchCommand(cxt).Command,
		NewDownloadCommand(cxt).Command,
	)

	return c
}
