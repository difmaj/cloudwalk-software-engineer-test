FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o cloudwalk-software-engineer-test ./cmd

CMD ["./cloudwalk-software-engineer-test"]
