package main

import (
	"os"

	"github.com/carolynvs/az-cli/cmd/az"
)

func main() {
	cmd := az.NewRootCommand().Command
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
