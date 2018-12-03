GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
DEP=dep ensure
BINARY_NAME=testurtle

deps:
	$(DEP)
test:
	$(GOTEST) -v ./...
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
