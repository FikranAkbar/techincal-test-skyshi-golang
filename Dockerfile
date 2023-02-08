FROM golang:1.19

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go mod tidy

RUN chmod +x /app/start.sh

CMD ["go run main.go"]