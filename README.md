# Validator Service & CLI
[![Go](https://img.shields.io/badge/Go-1.18-blue.svg)](https://golang.org)
[![build](https://github.com/arttet/validator-service/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/arttet/validator-service/actions/workflows/build.yml)
[![tests](https://github.com/arttet/validator-service/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/arttet/validator-service/actions/workflows/tests.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/arttet/validator-service/blob/main/LICENSE)

## Apply SQL migrations only

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```
```
goose -dir migrations mysql "username:password@tcp(127.0.0.1:3306)/test" status
goose -dir migrations mysql "username:password@tcp(127.0.0.1:3306)/test" up
goose -dir migrations mysql "username:password@tcp(127.0.0.1:3306)/test" status
```

## Build project
```sh
make all
```
