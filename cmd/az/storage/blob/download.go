package blob

import (
	"github.com/carolynvs/az-cli/cmd/az/command"
	"github.com/spf13/cobra"
)

// az storage blob download \
//  --container-name $AZURE_STORAGE_CONTAINER \
//  --name index.yaml \
//  --file index.yaml
type DownloadCommand struct {
	*command.Context
	*cobra.Command

	ContainerName string
	BlobName      string
	File string
}

func NewDownloadCommand(cxt *command.Context) *DownloadCommand {
	c := &DownloadCommand{
		Context: cxt,
	}

	c.Command = &cobra.Command{
		Use:     "download",
		Short:   "Downloads a blob to a file path",
		PreRunE: command.PreRunE(c),
		RunE:    command.RunE(c),
	}

	c.Flags().StringVarP(&c.ContainerName, "container-name", "c", "",
		"The container name")
	c.MarkFlagRequired("container-name")
	c.Flags().StringVarP(&c.File, "file", "f", "",
		"Path of file to write out to")
	c.MarkFlagRequired("file")
	c.Flags().StringVarP(&c.BlobName, "name", "n", "",
		"The blob name")
	c.MarkFlagRequired("name")

	return c
}

func (c *DownloadCommand) Validate(args []string) error {
	return nil
}

func (c *DownloadCommand) Run() error {
	return c.Download()
}

func (c *DownloadCommand) Download() error {
	return c.App.DownloadBlob(c.ContainerName, c.BlobName, c.File)
}
