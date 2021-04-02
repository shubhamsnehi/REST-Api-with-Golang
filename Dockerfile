FROM golang:1.15.8-alpine3.13

RUN mkdir /rest-api-app

ADD . /rest-api-app

WORKDIR /rest-api-app

RUN go build -o main

CMD ["/rest-api-app/main"]