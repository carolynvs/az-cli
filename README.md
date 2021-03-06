[![Build Status](https://travis-ci.org/carolynvs/az-cli.svg?branch=master)](https://travis-ci.org/carolynvs/az-cli)

This is a subset of the python Azure CLI, az, implemented in Go. The purpose
is to make it easier to install and use the Azure CLI in environments that don't
already have python installed (namely my CI build images).

I have only implemented/tested just enough to make this a useful tool for my own
purposes.

If there are az commands that you require, pull requests are welcome but
feature requests will be closed.

# Install
Replace `VERSION` in the URLs below with a [released
version](https://github.com/carolynvs/az-cli/releases), for example `v0.3.2`.

## Mac

```
curl -sLO https://github.com/carolynvs/az-cli/releases/download/VERSION/az-darwin-amd64
chmod +x az-darwin-amd64
mv az-darwin-amd64 /usr/local/bin/az
```

## Linux

```
curl -sLO https://github.com/carolynvs/az-cli/releases/download/VERSION/az-linux-amd64
chmod +x az-linux-amd64
mv az-linux-amd64 /usr/local/bin/az
```

## Dockerfile

```
RUN curl -sLo /usr/local/bin/az https://github.com/carolynvs/az-cli/releases/download/VERSION/az-linux-amd64 && \
    chmod +x /usr/local/bin/az
```

## Windows

The snippet below adds a directory to your PATH for the current session only.
You will need to find a permanent location for it and add it to your PATH.

```
mkdir -f ~\bin
iwr 'https://github.com/carolynvs/az-cli/releases/download/VERSION/az-windows-amd64.exe' -UseBasicParsing -OutFile ~\bin\az.exe
$env:PATH += ";~\bin"
```

# Supported Commands

Authentication is performed using the `AZURE_STORAGE_ACCOUNT` and `AZURE_STORAGE_ACCESS_KEY`, or `AZURE_STORAGE_CONNECTION_STRING` environment variables.

* `az --version`
* `az storage blob upload-batch --source --destination`
* `az storage blob download --container-name --name --file`
* `az storage container lease acquire -c -d`
* `az storage blob lease acquire -c -b -d`
