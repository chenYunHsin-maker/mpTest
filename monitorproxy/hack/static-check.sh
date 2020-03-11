#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

## Check coding style with specific directories excluded
#golint $(go list ./... | grep -vE "/vendor/|/pkg/api/|/pkg/apiclient/|/pkg/apimachinery/|/pkg/apiserver/")  > /tmp/a
## Filter out some style mistakes which we have to accept for now
#cat /tmp/a | grep -vE "generated|.pb.go|underscores|exported|ALL_CAPS|block ends with a return" | grep -vE "field|type|var|func|const .* should be" > /tmp/b || true
## Output messages left if any
#if [ -s /tmp/b ]; then
#    cat /tmp/b
#    exit 1
#fi
#echo "Success: golint passed."

# Check for shadowed variables, but ignore that of 'err' declaration 
go tool vet -shadow ./pkg ./cmd ./examples 2> /tmp/a || true
cat /tmp/a | grep -vE '"err" shadows declaration' > /tmp/b || true
if [ -s /tmp/b ]; then
    cat /tmp/b
    exit 1
fi
echo "Success: go vet -shadow passed."

# Check Go files for correctness 
go tool vet ./pkg ./cmd ./examples
echo "Success: go vet passed."
