FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go get github.com/swaggo/swag/gen@v1.16.3
RUN go get github.com/swaggo/http-swagger@v1.3.4
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3
RUN swag init -g cmd/main.go
RUN go build -o main.bin cmd/main.go

FROM alpine as release

WORKDIR /app

COPY --from=builder /app/main.bin /app/main.bin

EXPOSE 8080

ENTRYPOINT ["./main.bin"]
