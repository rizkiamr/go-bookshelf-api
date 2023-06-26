FROM golang:1.20 AS build

RUN mkdir -p /opt/build

WORKDIR /opt/build

# Copy all sources in
COPY . .

# Get dependencies for Go part of build
RUN go mod tidy

# Do the build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

FROM gcr.io/distroless/static
USER nobody:nobody
ENTRYPOINT []
WORKDIR /
COPY --from=build /opt/build/server .
ENV GIN_MODE=release