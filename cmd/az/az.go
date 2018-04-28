package az

import (
	"fmt"

	"github.com/carolynvs/az-cli/cmd/az/command"
	"github.com/carolynvs/az-cli/cmd/az/storage"
	"github.com/carolynvs/az-cli/pkg"
	"github.com/carolynvs/az-cli/pkg/az"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	*command.Context
	*cobra.Command

	Version bool
}

func NewRootCommand() *RootCommand {
	cxt := command.NewContext()
	c := &RootCommand{
		Context: cxt,
	}

	c.Command = &cobra.Command{
		Use:          "az",
		Short:        "Azure CLI, now with moar gophers!",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Enable tests to swap the output
			cxt.Output = cmd.OutOrStdout()

			app, err := az.NewApp()
			if err != nil {
				return err
			}
			cxt.App = app

			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if c.Version {
				c.PrintVersion(cxt)
				return nil
			}

			fmt.Fprint(cxt.Output, cmd.UsageString())
			return nil
		},
	}

	c.Flags().BoolVarP(&c.Version, "version", "v", false, "Show the application version")

	c.AddCommand(
		storage.NewSubCommand(cxt).Command,
	)

	return c
}

func (c *RootCommand) PrintVersion(cxt *command.Context) {
	fmt.Fprintf(cxt.Output, "az %s\n", pkg.Version)
}
