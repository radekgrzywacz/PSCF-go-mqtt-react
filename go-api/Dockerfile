FROM golang:1.24

WORKDIR /go-api
COPY . .

RUN go mod tidy
RUN go build -o main ./cmd/server

CMD ["./main"]
