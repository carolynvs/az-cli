package blob

import (
	"github.com/carolynvs/az-cli/cmd/az/command"
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
		NewUploadBatchCommand(cxt).Command,
	)

	return c
}
