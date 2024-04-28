FROM golang:alpine

WORKDIR /opt/app

COPY go.sum go.sum
COPY go.mod go.mod
RUN go mod download

COPY . .

EXPOSE 9095
EXPOSE 9096
