FROM golang:1.21.7-alpine

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -o main main.go

EXPOSE 8080

RUN chmod +x main

CMD ["./main"]