FROM golang:1.20-alpine AS build

# Install git
RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir -p /opt/build

WORKDIR /opt/build

# Copy all sources in
COPY . .

# Get dependencies for Go part of build
RUN go mod tidy

# This is a set of variables that the build script expects
ENV VERBOSE=0
ENV PKG=github.com/rizkiamr/go-bookshelf-api
ENV ARCH=amd64
ENV VERSION=test
ENV GIN_MODE=release

# Do the build
RUN go build -o server main.go

FROM alpine

USER nobody:nobody
COPY --from=build /opt/build/server /opt/server

CMD [ "/opt/server" ]