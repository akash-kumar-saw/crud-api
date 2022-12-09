FROM golang:1.20-rc-alpine3.17

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /crud-api

EXPOSE 8080

CMD ["/crud-api"]