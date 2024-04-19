FROM golang:latest

COPY ./ /app
WORKDIR /app

RUN go mod download
RUN go build -o main .

EXPOSE 8080

CMD ["go", "run", "main.go", "-port", "8000"]