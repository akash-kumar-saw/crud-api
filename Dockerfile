FROM golang:latest

WORKDIR /src

RUN export GO119MODULE=on

RUN cd /src && git clone https://github.com/akash-kumar-saw/crud-api

EXPOSE 8080