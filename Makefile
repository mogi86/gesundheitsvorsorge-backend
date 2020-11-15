.PHONY: build
build:
	go build -o server cmd/http/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: run
run: build
	./server

.PHONY: test
test:
	go test -v ./...