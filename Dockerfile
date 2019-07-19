FROM golang:1.12-stretch

COPY . /go/src/github.com/flga/freetype2
RUN go test -v -tags static github.com/flga/freetype2/2.10.1
RUN go test -v -tags 'static harfbuzz' github.com/flga/freetype2/2.10.1
RUN go test -v -tags 'static harfbuzz subset' github.com/flga/freetype2/2.10.1
