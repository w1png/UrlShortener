FROM golang:1.20.4-bullseye

WORKDIR /app
COPY . .

RUN go mod tidy

CMD ["go", "run", "server.go"]
