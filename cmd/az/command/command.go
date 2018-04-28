package command

import (
	"github.com/spf13/cobra"
)

// Command represents an svcat command.
type Command interface {
	// Validate and load the arguments passed to the command.
	Validate(args []string) error

	// Run a validated command.
	Run() error
}

// PreRunE validates and load the arguments passed to the command.
func PreRunE(cmd Command) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, args []string) error {
		return cmd.Validate(args)
	}
}

// RunE runs a validated command.
func RunE(cmd Command) func(*cobra.Command, []string) error {
	return func(_ *cobra.Command, args []string) error {
		return cmd.Run()
	}
}
