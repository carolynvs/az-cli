#!/usr/bin/env bash

set -xeuo pipefail

REPO_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
AZ=${AZ:-az}

echo "TEST: UploadBatch"
$REPO_ROOT/bin/$AZ storage blob upload-batch -s test/testdata/ -d azcli
