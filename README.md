# gesundheitsvorsorge-backend

- :warning: The development of this repository is working in progress.
- `gesundheitsvorsorge-backend` is backend side of application for management of health. 

## Requirement

- [golangci/golangci-lint](https://github.com/golangci/golangci-lint)
   - see: https://golangci-lint.run/usage/install/
   - e.g.) `curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0`

## Usage

- test

```bash
$ make test
```

- run

```bash
$ make run
```

- call api (need to use other terminal session)

```bash
$ curl -v -X GET "http://localhost:8000"
```

- stop

```bash
$ make stop
```
