GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.18","$(shell printf "$(GO_VERSION_SHORT)\n1.18" | sort -V | head -1)")
$(warning NEED GO VERSION >= 1.18. Found: $(GO_VERSION_SHORT))
endif

GITHUB_PATH=github.com/arttet/validator-service

###############################################################################

.PHONY: all
all: deps build

.PHONY: deps
deps: .deps-go

.PHONY: build
build:  .build

.PHONY: run
run:
	go run -v -race cmd/validator-cli/main.go --dsn "username:password@tcp(127.0.0.1:3306)/test"

.PHONY: test
test:
	go test -v -timeout 30s -coverprofile cover.out ./...
	go tool cover -func cover.out | grep -v -E '100.0%|total' || echo "OK"
	go tool cover -func cover.out | grep total | awk '{print ($$3)}'

.PHONY: bench
bench:
	go test -bench ./... -benchmem -cpuprofile cpu.out -memprofile mem.out -memprofilerate 1

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	find . -iname "*.go" | xargs gofmt -w

.PHONY: cover
cover:
	go tool cover -html cover.out

.PHONY: clean
clean:
	rm -rd ./bin/ || true

################################################################################

.PHONY: .deps-go
.deps-go:
	go mod download

################################################################################

.build: .build-validator-cli .build-validatord

.build-validator-cli: \
	$(eval SERVICE_NAME := validator-cli) \
	$(eval SERVICE_MAIN := cmd/$(SERVICE_NAME)/main.go) \
	$(eval SERVICE_EXE  := ./bin/$(SERVICE_NAME)) \
	.build-template

.build-validatord:
	GOOS=linux CGO_ENABLED=0 go build \
		-o ./bin/validatord cmd/validatord/main.go

.build-template:
	CGO_ENABLED=0 go build \
		-mod=mod \
		-o $(SERVICE_EXE)$(shell go env GOEXE) $(SERVICE_MAIN)

################################################################################
