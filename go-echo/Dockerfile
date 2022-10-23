# syntax=docker/dockerfile:1

## Build
FROM golang:alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -ldflags='-extldflags=-static -s' -tags osusergo,netgo -o /go-echo

## Deploy
FROM scratch
WORKDIR /
COPY --from=build /go-echo /go-echo
EXPOSE 8080
ENTRYPOINT ["/go-echo"]
