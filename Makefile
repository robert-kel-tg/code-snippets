NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

BINARY_NAME=main
BINARY_SRC=github.com/robertke/orders-service/${BINARY_NAME}
GO_LINKER_FLAGS=-ldflags "-s"

.PHONY: all clean deps build install

all: deps build

deps:
	git config --global http.https://gopkg.in.followRedirects true

	@echo "$(OK_COLOR)==> Installing glide dependencies$(NO_COLOR)"
	@go get -u github.com/Masterminds/glide
	@glide install

	@echo "$(OK_COLOR)==> Installing CompileDaemon$(NO_COLOR)"
	@go get github.com/githubnemo/CompileDaemon

build:
	@printf "$(OK_COLOR)==> Building binary$(NO_COLOR)\n"
	@GOOS=linux GOARCH=386 go build -o ${BINARY_NAME}

install:
	@printf "$(OK_COLOR)==> Installing binary$(NO_COLOR)\n"
	@go install -v ${BINARY_SRC}