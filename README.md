# gesundheitsvorsorge-backend

- :warning: The development of this repository is working in progress.
- `gesundheitsvorsorge-backend` is backend side of application for management of health. 

## Requirement

- [golangci/golangci-lint](https://github.com/golangci/golangci-lint)
   - see: https://golangci-lint.run/usage/install/
   - e.g.) `curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0`

## Usage

- create key for JWT

```bash
$ openssl genrsa 4096 > private.key
$ openssl rsa -pubout < private.key > public.key
```

- test

```bash
$ make test
```

- run

```bash
$ make run
```

- call api (need to use other terminal session)
    - see: [HOW_TO_EXECUTE](./doc/execute/HOW_TO_EXECUTE.md)

- stop

```bash
$ make stop
```
