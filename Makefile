BINARY = keyfob
BUILD_DIR = bin
TEST_REPORT = tests.xml

VERSION = "Make"

LDFLAGS = -ldflags "-X main.version=${VERSION}"

.PHONY: test fmt

#-----------------------------------------------------------------------
# Rules of Rules : Grouped rules that _doathing_
#-----------------------------------------------------------------------

all: clean fmt lint test build

#-----------------------------------------------------------------------
# Build
#-----------------------------------------------------------------------
build:
	mkdir ${BUILD_DIR} | true
	go build ${LDFLAGS} -o ${BUILD_DIR} ./...

#-----------------------------------------------------------------------
# Testing & Linting
#-----------------------------------------------------------------------

fmt:
	go fmt $$(go list ./...)

lint:
	golangci-lint run

test:
	go test -v ./... 2>&1 | go2xunit -output ${TEST_REPORT}

#-----------------------------------------------------------------------
# Utilities
#-----------------------------------------------------------------------
clean:
	rm -f ${TEST_REPORT}
	rm -rf ${BUILD_DIR}