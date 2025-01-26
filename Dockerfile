FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /account-service .

FROM alpine:latest
WORKDIR /
COPY --from=builder /account-service /account-service

EXPOSE 8080
CMD ["/account-service"]