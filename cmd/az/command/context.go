package command

import (
	"io"

	"github.com/carolynvs/az-cli/pkg/az"
)

// Context is ambient data necessary to run any az command.
type Context struct {
	// Output should be used instead of directly writing to stdout/stderr, to enable unit testing.
	Output io.Writer

	// svcat application, the library behind this cli
	App *az.App
}

// NewContext for the svcatt cli
func NewContext() *Context {
	return &Context{}
}
