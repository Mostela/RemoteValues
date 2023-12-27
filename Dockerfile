FROM golang:alpine as builder
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . .
RUN go build -o ./RemoteValues

ENV PORT=8080
ENV GIN_MODE=release

EXPOSE 8080

CMD ["./RemoteValues"]