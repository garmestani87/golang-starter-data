ARG GO_VERSION=1.23.2

FROM golang:${GO_VERSION}-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN GOPROXY=https://goproxy.io,direct go mod download

COPY . ./

RUN go build -v -o server ./main.go


# Start fresh from a smaller image
FROM alpine:latest
RUN apk add ca-certificates

COPY --from=builder /app/server /app/server
COPY --from=builder /app/config/application-dev.yml /app/config/application-dev.yml
COPY --from=builder /app/docs /app/docs


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/server"]