#!/bin/bash

set -e

STARTTIME=$(date +%s)
OS_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${OS_ROOT}/hack/common.sh"

GO_VERSION=($(go version))
echo "Detected go version: $(go version)"

echo "get go tools"
echo "get godep"
go get github.com/tools/godep
echo "get vet"
go get golang.org/x/tools/cmd/vet
echo "get cover"
go get golang.org/x/tools/cmd/cover 
echo "get tools finished."

# Check out a stable commit for go vet in order to version lock it to something we can work with
pushd $GOPATH/src/golang.org/x/tools >/dev/null 2>&1
  git fetch
  git checkout c262de870b618eed648983aa994b03bc04641c72 
popd >/dev/null 2>&1

# Re-install using this version of the tool
go install golang.org/x/tools/cmd/vet


ret=$?; ENDTIME=$(date +%s); echo "$0 took $(($ENDTIME - $STARTTIME)) seconds"; exit "$ret"
