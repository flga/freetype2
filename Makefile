
ifeq ("${OS}", "linux")
goLDFlags = -ldflags "-linkmode external -extldflags -static"
endif

test:
	CGO_ENABLED=1 GOARCH=$(ARCH) go test -v -tags 'static' $(goLDFlags) -coverage github.com/flga/freetype2/$(VERSION)
	CGO_ENABLED=1 GOARCH=$(ARCH) go test -v -tags 'static harfbuzz' $(goLDFlags) -coverage github.com/flga/freetype2/$(VERSION)
	CGO_ENABLED=1 GOARCH=$(ARCH) go test -v -tags 'static harfbuzz subset' $(goLDFlags) -coverage github.com/flga/freetype2/$(VERSION)