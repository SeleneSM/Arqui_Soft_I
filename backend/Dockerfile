
FROM golang:1.22.1

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /docker-go

EXPOSE 8090

CMD ["/docker-go"]
