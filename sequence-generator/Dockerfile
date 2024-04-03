FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY ./sequence-generator/*.go .

RUN go mod download && go build -o main .

EXPOSE 9090

CMD ["./main"]
