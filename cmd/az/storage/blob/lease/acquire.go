package lease

import (
	"fmt"

	"github.com/carolynvs/az-cli/cmd/az/command"
	"github.com/spf13/cobra"
)

// az storage blob lease acquire \
// --container-name $AZURE_STORAGE_CONTAINER \
// --blob-name index.yaml \
// --lease-duration 60
type AcquireCommand struct {
	*command.Context
	*cobra.Command

	ContainerName string
	BlobName      string
	Duration      int32
}

func NewAcquireCommand(cxt *command.Context) *AcquireCommand {
	c := &AcquireCommand{
		Context: cxt,
	}

	c.Command = &cobra.Command{
		Use:     "acquire",
		Short:   "Requests a new lease",
		PreRunE: command.PreRunE(c),
		RunE:    command.RunE(c),
	}

	c.Flags().StringVarP(&c.ContainerName, "container-name", "c", "",
		"The container name")
	c.MarkFlagRequired("container-name")
	c.Flags().StringVarP(&c.BlobName, "blob-name", "b", "",
		"The blob name")
	c.MarkFlagRequired("blob-name")
	c.Flags().Int32Var(&c.Duration, "lease-duration", -1,
		"Specifies the duration of the lease, in seconds, or negative one (-1) for a lease that never expires. A non-infinite lease can be between 15 and 60 seconds. A lease duration cannot be changed using renew or change. Default is -1 (infinite lease)")

	return c
}

func (c *AcquireCommand) Validate(args []string) error {

	return nil
}

func (c *AcquireCommand) Run() error {
	return c.Acquire()
}

func (c *AcquireCommand) Acquire() error {
	leaseID, err := c.App.AcquireBlobLease(c.ContainerName, c.BlobName, c.Duration)
	if err != nil {
		return err
	}

	fmt.Fprintln(c.Output, leaseID)
	return nil
}
