FROM golang:1.24.3-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

CMD ["./main"]
