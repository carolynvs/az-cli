package az

import (
	"net/http"

	"github.com/Azure/azure-storage-blob-go/2017-07-29/azblob"
	"github.com/pkg/errors"
)

func (a *App) AcquireBlobLease(containerName, blobName string, duration int32) (leaseID string, err error) {
	containerURL, err := a.buildContainerURL(containerName)
	if err != nil {
		return "", err
	}
	blobURL := containerURL.NewBlobURL(blobName)

	response, err := blobURL.AcquireLease(a.cxt, "", duration, azblob.HTTPAccessConditions{})
	if err != nil {
		return "", err
	}

	if response.StatusCode() != http.StatusCreated {
		return "", errors.New(response.Status())
	}

	return response.LeaseID(), nil
}
