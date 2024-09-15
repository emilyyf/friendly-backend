FROM golang:1.22.6 AS builder

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

EXPOSE 3000

CMD ["go", "run" , "cmd/main.go"]
