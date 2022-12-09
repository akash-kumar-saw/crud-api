FROM golang:latest

WORKDIR /src

RUN export GO111MODULE=on

RUN cd /src && git clone https://github.com/akash-kumar-saw/crud-api

RUN cd /src/crud-api/ && go build

EXPOSE 8080

ENTRYPOINT ["/src/crud-api/main"]
