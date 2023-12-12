.PHONY: build
build:
	go build

.PHONY: test
test:
	go test -coverpkg=./... ./...

.PHONY: release
release: build test
