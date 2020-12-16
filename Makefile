.PHONY: build
build:
	go build -o server cmd/http/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run --disable-all \
		-E govet \
		-E gocyclo \
		-E gofmt \
		-E goimports \
		-E misspell \
		-E wsl \
		./...

.PHONY: run-for-local
run-for-local: build
	./server

.PHONY: run
run: build
	docker-compose build
	docker-compose up -d
	make migration

.PHONY: stop
stop:
	docker-compose stop

.PHONY: test
test:
	go test -v ./...

.PHONY: migration
migration:
	go build -o migration \
		-ldflags "-X main.user=gesundheitsvorsorge -X main.password=gesundheitsvorsorge -X main.host=0.0.0.0 -X main.port=3306 -X main.DBName=gesundheitsvorsorge_db" \
		cmd/db/main.go
	./migration
