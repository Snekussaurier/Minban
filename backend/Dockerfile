# Build stage
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/. .
RUN go build -o main .

# Final stage
FROM alpine:latest

RUN apk add --no-cache sqlite

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 9916

CMD ["./main"]