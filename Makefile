#.SILENT:

LOCAL_BIN:=$(CURDIR)/bin
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.38.0

run:
	go run main.go


# install golangci-lint binary
.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG) && \
		go build -ldflags "-X 'main.version=$(GOLANGCI_TAG)' -X 'main.commit=test' -X 'main.date=test'" -o $(LOCAL_BIN)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint && \
		rm -rf $$tmp
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

# run diff lint like in pipeline
.PHONY: .lint
.lint: install-lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run  --config=.golangci.pipeline.yaml ./...

# golangci-lint diff master
.PHONY: lint
lint: .lint