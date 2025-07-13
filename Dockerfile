FROM golang:1.24.5 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . ./

RUN go build -o server ./cmd/server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/certs /app/certs

EXPOSE 8443

CMD ["./server"]
