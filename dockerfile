FROM golang:1.20.4-bullseye

WORKDIR /app
COPY . .

RUN go mod tidy
EXPOSE 8080

CMD ["go", "run", "."]
