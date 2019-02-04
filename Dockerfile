FROM golang:1.10-alpine

ENV GOPATH=/go/

RUN mkdir -p /go/src/github.com/sqmmm/event-service
WORKDIR /go/src/github.com/sqmmm/event-service

ADD . .

RUN go build -o events -i ./cmd/events/

CMD ["./events"]
