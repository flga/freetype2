#!/bin/sh
set -e

OS=$(go env GOHOSTOS)
if [[ "$OS" == "linux" ]]; then
    goLDFlags='-ldflags "-linkmode external -extldflags -static"'
fi

CGO_ENABLED=1 GOARCH=$(ARCH) go test -v -tags 'static' $(goLDFlags) -coverage github.com/flga/freetype2/$(VERSION)
CGO_ENABLED=1 GOARCH=$(ARCH) go test -v -tags 'static harfbuzz' $(goLDFlags) -coverage github.com/flga/freetype2/$(VERSION)
CGO_ENABLED=1 GOARCH=$(ARCH) go test -v -tags 'static harfbuzz subset' $(goLDFlags) -coverage github.com/flga/freetype2/$(VERSION)