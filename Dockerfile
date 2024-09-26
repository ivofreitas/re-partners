FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o main.bin cmd/main.go
FROM alpine as release

WORKDIR /app

COPY --from=builder /app/main.bin /app/main.bin

EXPOSE 8080

ENTRYPOINT ["./main.bin"]
