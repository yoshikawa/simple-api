FROM golang:latest

# Set go bin which doesn't appear to be set already.
ENV GOPATH /go
ENV GOBIN /go/bin

# build directories
ADD . $GOPATH/src/github.com/yoshikawataiki/simple-api
WORKDIR $GOPATH/src/github.com/yoshikawataiki/simple-api


RUN go get bitbucket.org/liamstask/goose/cmd/goose
# Go dep
RUN go get -u github.com/golang/dep/... && dep ensure -vendor-only