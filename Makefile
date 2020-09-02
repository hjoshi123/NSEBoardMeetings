GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test ./...

OS := $(shell uname -s | awk '{print tolower($$0)}')

TAG = $$(git rev-parse --short HEAD)

BINARY = main

GOARCH = amd64

LDFLAGS = -ldflags

.PHONY: run
run: bin #this will cause "bin" target to be build first
	./$(BINARY) # Execute the binary

.PHONY: bin
bin:
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=${GOARCH} go build ${LDFLAGS} "-w" -a -o ${BINARY} . ;

# Runs unit tests.
.PHONY: test
test:
	$(GOTEST)

# Generates a coverage report
.PHONY: cover
cover:
	${GOCMD} test -coverprofile=coverage.out ./... && ${GOCMD} tool cover -html=coverage.out

# Remove coverage report and the binary.
.SILENT: clean
.PHONY: clean
clean:
	$(GOCLEAN)
	@rm -f ${BINARY}
	@rm -f coverage.out

.PHONY: deps
deps:
	$(GOCMD) mod download
