FROM golang:alpine3.19

RUN apk update && \
    apk upgrade && \
    apk add bash
WORKDIR /go/src

COPY ./main.go /go/src

CMD ["/bin/bash"]
