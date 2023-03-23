GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=computer
GOMOD=$(GOCMD) mod

export GO111MODULE=on

all: deps test build
build: 
		$(GOGET)
		$(GOBUILD)
test:
		$(GOTEST) -v ./...

clean:
		$(GOCLEAN)
		rm -f go.sum
fclean:
		$(GOCLEAN)
		rm -f go.sum
		rm -f $(BINARY_NAME)

deps:
		$(GOMOD) tidy
		$(GOGET) computer/src/prompt

re: deps test build

.PHONY: all build test clean fclean