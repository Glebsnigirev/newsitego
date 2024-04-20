# FROM golang:latest

# COPY ./ /app
# WORKDIR /app

# RUN go mod vendor
# RUN go build -o main .

# EXPOSE 8080

# ENTRYPOINT ./main
                             
FROM golang:latest

COPY ./ /app
WORKDIR /app

RUN go mod download
CMD ["go", "run", "main.go"]
