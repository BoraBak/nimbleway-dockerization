FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY ./web-server/*.go .

RUN go mod download && go build -o main .

EXPOSE 8080

CMD ["./main"]
