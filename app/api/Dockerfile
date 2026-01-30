FROM golang:1.25-alpine as builder

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ../.. .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

FROM alpine:latest

WORKDIR /root/

ENV GIN_MODE=release
ENV PORT=8080

COPY --from=builder /app/main .
COPY --from=builder /app/docs .

EXPOSE 8080
CMD ["./main"]