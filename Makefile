RELEASE_BIN:=bin/file
GOLANGCI_LINT_VERSION:=v2.4.0

.PHONY: build
build: fmt
	CGO_ENABLED=0 go build -o $(RELEASE_BIN)

#####################################################################
# Rules for verification, formatting, linting, testing and cleaning #
#####################################################################

.PHONY: fmt
fmt: ## format golang source code files.
	go fmt ./...

.PHONY: lint
lint: golangci-lint-install
	CGO_ENABLED=0 $(GOLANGCILINT_BIN) run

#########################################
# Tools                                 #
#########################################

GOBIN=$(shell pwd)/bin
GOLANGCILINT_BIN := $(GOBIN)/golangci-lint-${GOLANGCI_LINT_VERSION}
.PHONY: golangci-lint-install
golangci-lint-install: ## Download golangci-lint locally if necessary.
	$(call go-install-tool,$(GOLANGCILINT_BIN),github.com/golangci/golangci-lint/v2/cmd/golangci-lint,${GOLANGCI_LINT_VERSION})


# go-install-tool will 'go install' any $1 package.
define go-install-tool
@[ -f $(1) ] || { \
set -e; \
echo "Downloading $(2)" ;\
version=@$(3) ;\
url=$(2) ;\
if echo $$version | grep 'v[2-9][0-9]*' -q; then \
	echo $$url | grep '/v[2-9][0-9]*/' -q || version="/$$(printf $$version | grep -o 'v[2-9][0-9]*')$$version" ;\
fi ;\
GOBIN=$(GOBIN) go install $(2)$$version ;\
binary=$$(echo "$$url" | rev | cut -d'/' -f 1 | rev) ;\
mv "$(GOBIN)/$${binary}" $(1) ;\
}
endef
