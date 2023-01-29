# syntax=docker/dockerfile:1

####
## BUILD
####

# Base image
FROM golang:1.19-alpine AS build

RUN addgroup -S webgroup && adduser -S webuser -G webgroup

WORKDIR /app

COPY ./go.mod .
COPY ./*.go .
#RUN go mod download

RUN pwd && ls -apl
RUN go build -o /homehelp

####
## DEPLOY
####

FROM alpine AS server

WORKDIR /

COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /homehelp /homehelp
# Run as non-root user
USER webuser

EXPOSE 8080

ENTRYPOINT ["/homehelp"]
