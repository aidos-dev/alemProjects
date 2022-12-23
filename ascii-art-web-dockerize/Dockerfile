# syntax=docker/dockerfile:1

##Build
FROM golang:1.16-alpine3.15 AS build
LABEL build=temp
WORKDIR /app
COPY ./ ./ 

RUN go build cmd/main.go

##Deploy
FROM alpine:3.15
RUN mkdir app
WORKDIR /app

LABEL version="1.0"
LABEL author="@Optimus"

COPY --from=build /app/main ./

COPY webFiles ./


EXPOSE 8080


ENTRYPOINT [ "./main" ]
