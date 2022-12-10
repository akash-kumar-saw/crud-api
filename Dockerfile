# Pulling the latest Golang Docker Image
FROM golang:latest

# Turning Go Module On
RUN export GO119MODULE=on

# Opening Port 8080 for Network Communication
EXPOSE 8080
