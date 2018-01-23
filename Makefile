NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

BINARY_NAME=api
REPO=github.com/robertke/orders-service/${BINARY_NAME}
BINARY_SRC=${REPO}
GO_LINKER_FLAGS=-ldflags "-s"
DIR_OUT=$(CURDIR)/out

.PHONY: all clean deps build install

all: deps

deps:
	git config --global http.https://gopkg.in.followRedirects true

	@echo "$(OK_COLOR)==> Installing glide dependencies$(NO_COLOR)"
	@go get -u github.com/Masterminds/glide
	@glide install

	@echo "$(OK_COLOR)==> Installing CompileDaemon$(NO_COLOR)"
	@go get github.com/githubnemo/CompileDaemon

build:
	@printf "$(OK_COLOR)==> Building binary$(NO_COLOR)\n"
	@go build -o ${DIR_OUT}/${BINARY} ${GO_LINKER_FLAGS} ${BINARY_SRC}

install:
	@printf "$(OK_COLOR)==> Installing binary$(NO_COLOR)\n"
	@go install -v ${BINARY_SRC}