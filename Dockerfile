
FROM golang:1.16-alpine
WORKDIR /app
COPY . .
RUN go mod download
COPY .env .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

