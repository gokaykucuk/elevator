# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=runner
BINARY_UNIX=$(BINARY_NAME)_unix

test:
		$(GOTEST) -v ./...
ci_deps:
		$(GOGET) -u golang.org/x/tools/cmd/cover
		$(GOGET) -u github.com/mattn/goveralls

# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
# docker-build:
#         docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v