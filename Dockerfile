FROM golang:1.9-alpine
MAINTAINER kel.robert@gmail.com

ENV APP_DIR $GOPATH/src/github.com/robertke/orders-service

# Copy the local package files to the container’s workspace.
ADD . ${APP_DIR}

# Set workng dir
WORKDIR ${APP_DIR}

EXPOSE 8080

RUN apk add --update git curl make && \
    rm -rf /var/cache/apk/*

RUN make