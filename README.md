# Availability Check Service

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
