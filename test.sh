#!/bin/bash
set -e

OS=$(go env GOHOSTOS)
if [ "$OS" == "linux" ]; then
    goLDFlags=(-ldflags "-linkmode external -extldflags -static")
fi

cd $VERSION

CGO_ENABLED=1 GOARCH=$ARCH go test -tags 'static' "${goLDFlags[@]}" -cover
CGO_ENABLED=1 GOARCH=$ARCH go test -tags 'static harfbuzz' "${goLDFlags[@]}" -cover
CGO_ENABLED=1 GOARCH=$ARCH go test -tags 'static harfbuzz subset' "${goLDFlags[@]}" -cover