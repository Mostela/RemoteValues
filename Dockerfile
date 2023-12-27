FROM golang:alpine as builder
WORKDIR app
ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go build -o ./keyRotationK8S

ENV PORT=8080
ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ./keyRotationK8S