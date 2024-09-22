FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go install github.com/air-verse/air

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 3000

CMD ["air"]
