.PHONY: build
build:
	go build -o server cmd/http/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: run-for-local
run-for-local: build
	./server

.PHONY: run
run: build
	docker-compose build
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop

.PHONY: test
test:
	go test -v ./...