FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go get github.com/gin-contrib/cors 

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
