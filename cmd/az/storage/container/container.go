package container

import (
	"github.com/carolynvs/az-cli/cmd/az/command"
	"github.com/carolynvs/az-cli/cmd/az/storage/container/lease"
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
			Use:   "container",
			Short: "container commands",
		},
	}

	c.AddCommand(
		lease.NewSubCommand(cxt).Command,
	)

	return c
}
