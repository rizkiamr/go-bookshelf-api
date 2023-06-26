FROM golang:1.20 AS build

RUN mkdir -p /opt/build

WORKDIR /opt/build

# Copy only necessary files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the files
COPY . .

# Do the build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

FROM gcr.io/distroless/static
USER nobody:nobody
WORKDIR /
COPY --from=build /opt/build/server .
COPY app.env .
ENV GIN_MODE=release
ENTRYPOINT []