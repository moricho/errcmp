GOLANGCI_LINT_VERSION=v1.23.1
GOLANGCI_LINT_COMMAND := $(shell which golangci-lint)

install_lint:
	go get -u golang.org/x/lint/golint
	go get -u github.com/fatih/errwrap
ifndef GOLANGCI_LINT_COMMAND
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)
endif

install_test:
	curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > $(CMD_PATH)/cc-test-reporter
	chmod +x $(CMD_PATH)//cc-test-reporter

lint:
	go vet ./...
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status=true
	golangci-lint run ./...
	errwrap ./...

test:
	go test -parallel 3 -count 1 ./... -cover

goimports:
	find -E . -type f -iregex ".*\.(go)" | xargs -I _ goimports -local "github.com/moricho" -w _
