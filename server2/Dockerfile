FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8002

CMD ["./main"]