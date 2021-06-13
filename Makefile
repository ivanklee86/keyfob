TEST_REPORT = tests.xml

test:
	godep go test -v ./... 2>&1 | go2xunit -output ${TEST_REPORT} ;

.PHONY: test