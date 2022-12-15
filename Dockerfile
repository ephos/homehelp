# syntax=docker/dockerfile:1

# Base image
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./ 

RUN go mod download

COPY *.go ./ 

RUN go build -o /homehelp

EXPOSE 8080

CMD [ "/homehelp" ]
