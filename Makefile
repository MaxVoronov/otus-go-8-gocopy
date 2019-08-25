BINARY=gocopy
ARCHITECTURES=386 amd64
PLATFORMS=darwin linux windows

help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## build: Build binary into directory "bin"
build:
	@go build -o ./bin/$(BINARY)

## build-multi: Build binaries for different platforms and architectures into directory "bin"
build-multi:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES),\
	$(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -o ./bin/$(BINARY)-$(GOOS)-$(GOARCH))))

## install: Build and install
install:
	@go install

## lint: Lint project code (using golangci-lint)
lint:
	@golangci-lint run -v

## test: Start project testing
test:
	@go test -v ./...

## cover: Generate test coverage report and open it
cover:
	@go test -v -coverprofile=/tmp/$(BINARY)-cover.out ./...
	@go tool cover -html=/tmp/$(BINARY)-cover.out