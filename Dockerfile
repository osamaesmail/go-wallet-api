FROM golang:1.18.5-buster AS builder

RUN apt-get update && apt-get install -y git
RUN go install github.com/cespare/reflex@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN cp /go/bin/reflex /usr/bin/reflex
RUN cp /go/bin/dlv /usr/bin/dlv

WORKDIR /app

ENV GO111MODULE on

COPY go.mod .
COPY go.sum .
COPY .env .

RUN go mod download
