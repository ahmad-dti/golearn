# build stage with module caching
FROM golang:1.21-alpine AS builder

# install git for modules (if you depend on private repos)
RUN apk add --no-cache git

WORKDIR /app

# copy go.mod first to leverage layer cache
COPY go.mod .
# generate go.sum and download modules
RUN go mod tidy

# then copy rest of source and build
COPY . ./
# build only the server command
RUN go build -o /go/bin/app -v ./cmd/server

# final minimal runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
# copy the compiled binary into the working directory
COPY --from=builder /go/bin/app ./app

EXPOSE 8080
ENTRYPOINT ["/app/app"]
LABEL Name=golearn Version=0.0.1
