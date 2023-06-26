FROM golang:1.20 AS build

RUN mkdir -p /opt/build

WORKDIR /opt/build

# Copy all sources in
COPY . .

# Get dependencies for Go part of build
RUN go mod tidy

# Do the build
RUN go build -o server main.go

FROM alpine:3 as alpine
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates

FROM gcr.io/distroless/static
USER nobody:nobody
COPY --from=alpine /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=alpine /etc/passwd /etc/passwd
COPY --from=build /opt/build/server /opt/server
ENV GIN_MODE=release

CMD [ "/opt/server" ]