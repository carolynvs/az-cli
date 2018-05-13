# Azure Go CLI
This is a subset of the python Azure CLI, az, implemented in Go. The purpose
is to make it easier to install and use the Azure CLI in environments that don't
already have python installed (namely my CI build images).

I have only implemented/tested just enough to make this a useful tool for my own
purposes but would love it if the Azure team was interested in adopting it and
implementing the full set of commands.

If there are az commands that you require, pull requests are welcome but
feature requests will be closed.

# Install

## Mac

```
curl -sLO https://github.com/carolynvs/az-cli/releases/download/latest/az-darwin-amd64
chmod +x az-darwin-amd64
mv az-darwin-amd64 /usr/local/bin/az
```

## Linux

```
curl -sLO https://github.com/carolynvs/az-cli/releases/download/latest/az-linux-amd64
chmod +x az-linux-amd64
mv az-linux-amd64 /usr/local/bin/az
```

## Dockerfile
I recommend that you replace `latest` in the URL below with a specific release, e.g. `v0.1.0`.

```
RUN curl -sLo /usr/local/bin/az https://github.com/carolynvs/az-cli/releases/download/latest/az-linux-amd64 && \
    chmod +x /usr/local/bin/az
```

## Windows

The snippet below adds a directory to your PATH for the current session only.
You will need to find a permanent location for it and add it to your PATH.

```
mkdir -f ~\bin
iwr 'https://github.com/carolynvs/az-cli/releases/download/latest/az-windows-amd64.exe' -UseBasicParsing -OutFile ~\bin\az.exe
$env:PATH += ";~\bin"
```

# Supported Commands

Authentication is performed using the `AZURE_STORAGE_ACCOUNT` and `AZURE_STORAGE_ACCESS_KEY` environment variables.

* `az --version`
* `az storage blob upload-batch --source --destination`