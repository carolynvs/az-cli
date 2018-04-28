package az

import (
	"context"
	"errors"
	"os"

	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/2017-07-29/azblob"
)

type App struct {
	cxt        context.Context
	Credential *azblob.SharedKeyCredential
	Pipeline   pipeline.Pipeline
}

func NewApp() (*App, error) {
	app := &App{
		cxt: context.Background(),
	}
	err := app.loadCredentials()
	return app, err
}

func (a *App) loadCredentials() error {
	// From the Azure portal, get your storage account name and key and set environment variables.
	accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY")
	if accountName == "" || accountKey == "" {
		return errors.New("AZURE_STORAGE_ACCOUNT and AZURE_STORAGE_ACCESS_KEY must be set")
	}

	a.Credential = azblob.NewSharedKeyCredential(accountName, accountKey)
	a.Pipeline = azblob.NewPipeline(a.Credential, azblob.PipelineOptions{})

	return nil
}
