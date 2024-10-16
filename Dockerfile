FROM golang:alpine3.17 AS builder

WORKDIR /go/src/app
COPY  . . 
RUN go mod download

RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/channel ./server.go

FROM alpine:3.17
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=builder /go/src/app/bin /go/bin
COPY --from=builder /go/src/app/.env .
EXPOSE 8080
CMD ["/go/bin/channel", "--port", "8080"]
