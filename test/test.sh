#!/usr/bin/env bash

set -euo pipefail

REPO_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
AZ=${AZ:-az}

echo "TEST: UploadBatch"
$REPO_ROOT/bin/$AZ storage blob upload-batch -s test/testdata -d azcli

echo "TEST: DownloadBlob"
$REPO_ROOT/bin/$AZ storage blob download --container-name azcli --name a.txt --file test/testdata/a.txt

echo "TEST: AcquireContainerLease"
$REPO_ROOT/bin/$AZ storage container lease acquire --container-name azcli --lease-duration 15

echo "TEST: AcquireBlobLease"
$REPO_ROOT/bin/$AZ storage blob lease acquire --container-name azcli --blob-name a.txt --lease-duration 15
