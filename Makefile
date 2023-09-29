OS_NAME := $(shell uname)
ifeq ($(OS_NAME), Darwin)
OPEN := open
else
OPEN := xdg-open
endif

qa: analyze test

analyze:
	@go vet ./...
	@go run honnef.co/go/tools/cmd/staticcheck@latest --checks=all ./...

test:
	@go test -failfast -cover ./...

.PHONY: analyze \
				qa \
				test
