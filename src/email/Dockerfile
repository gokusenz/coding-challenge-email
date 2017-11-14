FROM golang:1.9-alpine

ADD . /go/src/web
RUN go install web
RUN rm -rf /go/src/web
ENTRYPOINT /go/bin/web

EXPOSE 80