SHELL := /bin/sh

TARGET := $(shell echo "bin/NSEBoardMeetings")
.DEFAULT_GOAL: $(TARGET)

PORT := 5000
export PORT

VERSION := 0.0.1
BUILD := `git rev-parse HEAD`

SRC = $(shell find . -type f -name '*.go' -not -path "./bin/*")

# making sure no name collision
.PHONY: all clean build uninstall fmt simplify run

all: uninstall build

$(TARGET): $(SRC)
	@go build -o $(TARGET)

clean:
	@rm -f $(TARGET)

build:
	@go build -o $(TARGET)

uninstall: clean
	@rm -f $$(which ${TARGET})

format:
	@gofmt -w $(SRC)

simplify:
	@gofmt -s -w $(SRC)

run: build
	@$(TARGET)%