FROM golang:1-23-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o api-gateway ./main

FROM alpine:3.19

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/api-gateway /api-gateway

EXPOSE 8080

CMD ["/api-gateway"]