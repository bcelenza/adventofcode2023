.PHONY: build
build:
	go build

.PHONY: test
test:
	go test -v -coverpkg=./... ./...

.PHONY: release
release: build test
