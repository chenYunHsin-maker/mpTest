#!/bin/bash

set -o errexit
set -o pipefail

if [[ -z $1 ]]; then
    echo "no work dir specified"
    exit 1
fi

WORK_DIR=$1

# Generate go files from protobuf files
cd $WORK_DIR/proto || exit 1
protoc --go_out=plugins=grpc:../ *.proto

# Inject custom tags to generated go files
cd ../
for i in *.pb.go; do
    [[ -f $i ]] || break
    protoc-go-inject-tag -input=$i
done

# Remove omitempty tag from generated json tags
ls *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
