package blob

import (
	"github.com/carolynvs/az-cli/cmd/az/command"
	"github.com/spf13/cobra"
)

// az storage blob upload \
//  --container-name $AZURE_STORAGE_CONTAINER \
//  --name index.yaml \
//  --file index.yaml
type UploadCommand struct {
	*command.Context
	*cobra.Command

	ContainerName string
	BlobName      string
	File          string
}

func NewUploadCommand(cxt *command.Context) *UploadCommand {
	c := &UploadCommand{
		Context: cxt,
	}

	c.Command = &cobra.Command{
		Use:     "upload",
		Short:   "Upload a file to a storage blob",
		PreRunE: command.PreRunE(c),
		RunE:    command.RunE(c),
	}

	c.Flags().StringVarP(&c.ContainerName, "container-name", "c", "",
		"The container name")
	c.MarkFlagRequired("container-name")
	c.Flags().StringVarP(&c.File, "file", "f", "",
		"Path of the file to upload as the blob content")
	c.MarkFlagRequired("file")
	c.Flags().StringVarP(&c.BlobName, "name", "n", "",
		"The blob name")
	c.MarkFlagRequired("name")

	return c
}

func (c *UploadCommand) Validate(args []string) error {
	return nil
}

func (c *UploadCommand) Run() error {
	return c.Upload()
}

func (c *UploadCommand) Upload() error {
	return c.App.UploadBlob(c.ContainerName, c.BlobName, c.File)
}
