package az

import (
	"context"
	"os"
	"regexp"

	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/2017-07-29/azblob"
	"github.com/pkg/errors"
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
	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT")
	accountKey := os.Getenv("AZURE_STORAGE_ACCESS_KEY")

	if accountName == "" || accountKey == "" {
		connString := os.Getenv("AZURE_STORAGE_CONNECTION_STRING")
		if connString == "" {
			errors.New("AZURE_STORAGE_ACCOUNT and AZURE_STORAGE_ACCESS_KEY or AZURE_STORAGE_CONNECTION_STRING must be set")
		}

		var err error
		accountName, accountKey, err = parseConnectionString(connString)
		if err != nil {
			return err
		}
	}

	a.Credential = azblob.NewSharedKeyCredential(accountName, accountKey)
	a.Pipeline = azblob.NewPipeline(a.Credential, azblob.PipelineOptions{})

	return nil
}

func parseConnectionString(connString string) (name string, key string, err error) {
	keyRegex := regexp.MustCompile("AccountKey=([^;]+)")
	keyMatch := keyRegex.FindAllStringSubmatch(connString, -1)

	nameRegex := regexp.MustCompile("AccountName=([^;]+)")
	nameMatch := nameRegex.FindAllStringSubmatch(connString, -1)

	if len(nameMatch) == 0 || len(keyMatch) == 0 {
		return "", "", errors.New("unexpected format for AZURE_STORAGE_CONNECTION_STRING, could not find AccountName=NAME and AccountKey=KEY in it")
	}

	accountKey := keyMatch[0][1]
	accountName := nameMatch[0][1]
	return accountName, accountKey, nil
}
