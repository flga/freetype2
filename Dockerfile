FROM golang:1.12-stretch

COPY . /go/src/github.com/flga/freetype2
RUN CGO_ENABLED=1 GOARCH=amd64 go test -tags static -ldflags "-linkmode external -extldflags -static" github.com/flga/freetype2/2.10.1
RUN CGO_ENABLED=1 GOARCH=amd64 go test -tags 'static harfbuzz' -ldflags "-linkmode external -extldflags -static" github.com/flga/freetype2/2.10.1
RUN CGO_ENABLED=1 GOARCH=amd64 go test -tags 'static harfbuzz subset' -ldflags "-linkmode external -extldflags -static" github.com/flga/freetype2/2.10.1
