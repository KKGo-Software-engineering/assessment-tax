# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' Makefile | column -t -s ':' | sed -e 's/^/ /'

## init: initial dependencies
.PHONY: init
init:
	@echo '== init dependencies =='
	go mod tidy
	go install github.com/joho/godotenv/cmd/godotenv@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	go install go.uber.org/mock/mockgen@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest



# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## lint: run golangci-lint and fix found issues
.PHONY: lint
lint:
	@echo '== linter =='
	golangci-lint run -v --fix

## audit: run quality control checks
.PHONY: audit
audit:
	@echo '== quality control =='
	go mod verify
	go vet ./...
	staticcheck -checks=all,-ST1000,-U1000 ./...
	govulncheck ./...

## gocyclo: run gocyclo for calculates cyclomatic complexities
.PHONY: gocyclo
gocyclo:
	@echo '== calculate cyclomatic complexities =='
	gocyclo .

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## mock: mockgen
.PHONY: mock
mock:
	@echo '== mock generate =='
	mockgen -source=internal/core/port/repository.go -destination=mocks/repository.go -package=mocks
	mockgen -source=internal/core/port/service.go -destination=mocks/service.go -package=mocks

## test: run all tests
.PHONY: test
test:
	@echo '== test application =='
	go test -v -race ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	@echo '== test application with coverage =='
	go test -v -race -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out
	go tool cover -html=coverage.out -o coverage.html

## run: run application
.PHONY: run
run:
	@echo '== run application =='
	godotenv -f .env go run main.go

## build: build application
.PHONY: build
build:
	@echo '== build application =='
	go mod tidy && \
	go mod vendor && \
	go build -mod=vendor -a -o ./out/main .

## clean: clean application
.PHONY: clean
clean:
	@echo '== clean application =='
	@rm -rf ./out ./vendor
