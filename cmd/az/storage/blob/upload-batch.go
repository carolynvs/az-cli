package blob

import (
	"github.com/carolynvs/az-cli/cmd/az/command"
	"github.com/spf13/cobra"
)

// az storage blob upload-batch --destination $AZURE_STORAGE_CONTAINER --source .
type UploadBatchCommand struct {
	*command.Context
	*cobra.Command

	Destination string
	Source      string
}

func NewUploadBatchCommand(cxt *command.Context) *UploadBatchCommand {
	c := &UploadBatchCommand{
		Context: cxt,
	}

	c.Command = &cobra.Command{
		Use:     "upload-batch",
		Short:   "Upload files from a local directory to a blob container",
		PreRunE: command.PreRunE(c),
		RunE:    command.RunE(c),
	}

	c.Flags().StringVarP(&c.Destination, "destination", "d", "",
		"The blob container where the files will be uploaded")
	c.MarkFlagRequired("destination")
	c.Flags().StringVarP(&c.Source, "source", "s", "",
		"The directory where the files to be uploaded are located")
	c.MarkFlagRequired("source")

	return c
}

func (c *UploadBatchCommand) Validate(args []string) error {
	return nil
}

func (c *UploadBatchCommand) Run() error {
	return c.UploadBatch()
}

func (c *UploadBatchCommand) UploadBatch() error {
	return c.App.UploadBatch(c.Source, c.Destination)
}
