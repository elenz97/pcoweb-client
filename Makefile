##@ Help

.SILENT: help
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build

.SILENT: build-linux
.PHONY: build-linux
build-linux: ## Build binary for Linux
	@CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -extldflags '-static'" -a -o pcoweb-client

.SILENT: build-armv7
.PHONY: build-armv7
build-armv7: ## Build binary for ARMv7 architecture
	@GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -ldflags="-s -w" -o pcoweb-client-armv7

.SILENT: build-arm64
.PHONY: build-arm64
build-arm64: ## Build binary for ARM64 architecture
	@GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -o pcoweb-client-arm64
