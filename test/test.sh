#!/usr/bin/env bash

set -euo pipefail

REPO_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
AZ=${AZ:-az}

echo "TEST: UploadBlob"
$REPO_ROOT/bin/$AZ storage blob upload -f test/testdata/a.txt -c azcli -n a.txt

echo "TEST: UploadBatch from reldir"
$REPO_ROOT/bin/$AZ storage blob upload-batch -s test/testdata -d azcli

echo "TEST: UploadBatch from curdir"
pushd test/testdata
$REPO_ROOT/bin/$AZ storage blob upload-batch -s . -d azcli
popd

echo "TEST: DownloadBlob"
$REPO_ROOT/bin/$AZ storage blob download --container-name azcli --name a.txt --file test/testdata/a.txt

echo "TEST: AcquireContainerLease"
$REPO_ROOT/bin/$AZ storage container lease acquire --container-name azcli --lease-duration 15

echo "TEST: AcquireBlobLease"
$REPO_ROOT/bin/$AZ storage blob lease acquire --container-name azcli --blob-name a.txt --lease-duration 15
