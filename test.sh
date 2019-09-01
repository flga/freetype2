#!/bin/bash
set -e

OS=$(go env GOHOSTOS)
if [ "$OS" == "linux" ]; then
    goLDFlags=(-ldflags "-linkmode external -extldflags -static")
fi

cd $VERSION

GODEBUG=cgocheck=2 CGO_ENABLED=1 GOARCH=$ARCH go test -tags 'static' "${goLDFlags[@]}" -cover
GODEBUG=cgocheck=2 CGO_ENABLED=1 GOARCH=$ARCH go test -tags 'static harfbuzz' "${goLDFlags[@]}" -cover
GODEBUG=cgocheck=2 CGO_ENABLED=1 GOARCH=$ARCH go test -tags 'static harfbuzz subset' "${goLDFlags[@]}" -cover

