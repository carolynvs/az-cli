package storage

import (
	"github.com/carolynvs/az-cli/cmd/az/command"
	"github.com/carolynvs/az-cli/cmd/az/storage/blob"
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
			Use:   "storage",
			Short: "storage commands",
		},
	}

	c.AddCommand(
		blob.NewSubCommand(cxt).Command,
	)

	return c
}
