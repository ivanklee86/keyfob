BINARY = keyfob
BUILD_DIR = bin
TEST_REPORT = tests.xml
COVERAGE_RAW = cover.out
COVERAGE_REPORT = cover.html

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
install:
	go get ./...

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
	go test --coverprofile ${COVERAGE_RAW} -v ./... 2>&1  | go-junit-report > report.xml
	go tool cover -html=${COVERAGE_RAW} -o ${COVERAGE_REPORT}

#-----------------------------------------------------------------------
# Utilities
#-----------------------------------------------------------------------
clean:
	rm -f ${TEST_REPORT}
	rm -rf ${BUILD_DIR}
	rm -f ${COVERAGE_RAW}
	rm -f ${COVERAGE_REPORT}