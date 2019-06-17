package az

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parseCredentials(t *testing.T) {
	testcases := []struct {
		name       string
		connString string
		err        string
	}{
		{name: "double semis", connString: "AccountName=myacct;AccountKey=mykey;"},
		{name: "trailing acct key", connString: "AccountName=myacct;AccountKey=mykey"},
		{name: "trailing acct name", connString: "AccountKey=mykey;AccountName=myacct"},
		{name: "full conn string", connString: "AccountName=myacct;DefaultEndpointsProtocol=https;AccountKey=mykey;EndpointSuffix=core.windows.net"},
		{name: "missing acct key", connString: "AccountName=myacct;", err: "unexpected format for AZURE_STORAGE_CONNECTION_STRING"},
		{name: "missing acct name", connString: "AccountKey=mykey;", err: "unexpected format for AZURE_STORAGE_CONNECTION_STRING"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			acctName, acctKey, err := parseConnectionString(tc.connString)

			if tc.err == "" {
				require.NoError(t, err)
				assert.Equal(t, "myacct", acctName)
				assert.Equal(t, "mykey", acctKey)
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.err)
			}
		})
	}
}

func Test_loadCredentials_ConnString(t *testing.T) {
	os.Setenv("AZURE_STORAGE_CONNECTION_STRING", "DefaultEndpointsProtocol=https;AccountName=myacct;AccountKey=fPzW0eoOZXxBOkSyOQdlx47eDdu8eWOESzwBlhyc0hVTCjkgcWPG4hNLEsg5aPWGVjeKBAOR98nDEL2sKqFAQg==;EndpointSuffix=core.windows.net")
	defer os.Unsetenv("AZURE_STORAGE_CONNECTION_STRING")

	a := &App{}
	err := a.loadCredentials()
	require.NoError(t, err, "load credentials failed")

	assert.Equal(t, "myacct", a.Credential.AccountName())
}

func Test_loadCredentials_NameAndKey(t *testing.T) {
	os.Setenv("AZURE_STORAGE_ACCOUNT", "myacct")
	os.Setenv("AZURE_STORAGE_ACCESS_KEY", "fPzW0eoOZXxBOkSyOQdlx47eDdu8eWOESzwBlhyc0hVTCjkgcWPG4hNLEsg5aPWGVjeKBAOR98nDEL2sKqFAQg==")
	defer func() {
		os.Unsetenv("AZURE_STORAGE_ACCOUNT")
		os.Unsetenv("AZURE_STORAGE_ACCESS_KEY")
	}()

	a := &App{}
	err := a.loadCredentials()
	require.NoError(t, err, "load credentials failed")

	assert.Equal(t, "myacct", a.Credential.AccountName())
}
