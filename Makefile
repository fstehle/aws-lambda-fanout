SHELL := /bin/bash

# The destination for binaries
BUILD_DIR = build

# The name of the executable (default is current directory name)
TARGET := ./src/FanOutConfigurator/FanOutConfigurator
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION :=0.9.0-rc1

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) "

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -path "./src/FanOutConfigurator/*")

.PHONY: all build clean install uninstall fmt simplify check run

all: check compile

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

compile: check goxcompile

goxcompile: export CGO_ENABLED=0
goxcompile: dependencies
	gox -arch amd64 -os darwin -os linux -os windows -output "$(BUILD_DIR)/{{.OS}}/$(NAME)/${TARGET}" ./src/FanOutConfigurator

clean:
	@rm -f $(TARGET)
	@rm -rf $(BUILD_DIR)

install: dependencies
	@go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

package: compile
	@for d in $$(ls build); do tar -czf $(BUILD_DIR)/${TARGET}-$${d}.tar.gz -C $(BUILD_DIR)/$${d} .; done

check: dependencies
	echo 'check start'
	@test -z $(shell gofmt -l ${TARGET}.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	#@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}

run: install
	@$(TARGET)

dependencies:
	go get github.com/aws/aws-sdk-go/aws/session
	go get gopkg.in/yaml.v2
	go get gopkg.in/fatih/set.v0
	echo 'dependencies done'


#export GOPATH=/home/schuenemann/project/MOIA/workplace/aws-lambda-fanout
#go get github.com/aws/aws-sdk-go/aws/session
#go get gopkg.in/yaml.v2
#go get gopkg.in/fatih/set.v0
#go build src/FanOutConfigurator/FanOutConfigurator.go