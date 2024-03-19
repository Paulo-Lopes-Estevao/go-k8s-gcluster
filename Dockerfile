FROM golang:1.21.7-alpine

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o main main.go

EXPOSE 8080

CMD ["./main"]