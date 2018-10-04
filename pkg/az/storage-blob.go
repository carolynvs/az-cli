package az

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/Azure/azure-storage-blob-go/2017-07-29/azblob"
	"github.com/pkg/errors"
)

func (a *App) UploadBlob(containerName, blobName, source string) error {
	containerURL, err := a.buildContainerURL(containerName)
	if err != nil {
		return err
	}

	return a.uploadFile(source, blobName, containerURL)
}

func (a *App) UploadBlobBatch(sourceDirectory, containerName string) error {
	containerURL, err := a.buildContainerURL(containerName)
	if err != nil {
		return err
	}

	numCores := runtime.GOMAXPROCS(0)
	files := make(chan string, 10*numCores)

	var wg sync.WaitGroup
	for i := 0; i < numCores; i++ {
		wg.Add(1)
		go func() {
			for path := range files {
				name, err := filepath.Rel(sourceDirectory, path)
				if err != nil {
					log.Printf("skipping %q: %s", path, err)
				}

				err = a.uploadFile(path, name, containerURL)
				if err != nil {
					log.Print(err)
				}
			}
			wg.Done()
		}()
	}

	err = filepath.Walk(sourceDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		files <- path
		return nil
	})
	close(files)

	wg.Wait()

	return err
}

func (a *App) uploadFile(path, name string, containerURL azblob.ContainerURL) error {
	blobURL := containerURL.NewBlockBlobURL(name)
	file, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "skipping file upload, could not open %s", path)
	}
	defer file.Close()

	_, err = azblob.UploadFileToBlockBlob(a.cxt, file, blobURL, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16})
	fmt.Printf("uploaded %s to %s\n", path, blobURL)

	return err
}

func (a *App) DownloadBlob(containerName, blobName, destination string) error {
	containerURL, err := a.buildContainerURL(containerName)
	if err != nil {
		return err
	}
	blobURL := containerURL.NewBlobURL(blobName)

	file, err := os.Create(destination)
	if err != nil {
		return errors.Wrapf(err, "cannot write to %s", destination)
	}

	return azblob.DownloadBlobToFile(a.cxt, blobURL, 0, 0, azblob.BlobAccessConditions{}, file, azblob.DownloadFromBlobOptions{})
}

func (a *App) buildContainerURL(containerName string) (azblob.ContainerURL, error) {
	rawURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", a.Credential.AccountName(), containerName)
	URL, err := url.Parse(rawURL)
	if err != nil {
		return azblob.ContainerURL{}, errors.Wrapf(err, "could not parse container URL %s", rawURL)
	}

	return azblob.NewContainerURL(*URL, a.Pipeline), nil
}
