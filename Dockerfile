FROM golang:1.9-alpine

RUN apk add --update curl git tzdata && \
    rm -rf /var/cache/apk/* && \
    curl https://glide.sh/get | sh && \
    cp -r -f /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    apk del tzdata

WORKDIR /go/src/email

COPY ./src/email .

RUN glide install && \
    go install .

CMD ["email"]
