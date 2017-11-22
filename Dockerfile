FROM golang:alpine


WORKDIR /go/src/github.com/reivaj05/micro-gen

ADD . /go/src/github.com/reivaj05/micro-gen

RUN apk -U add make git bash wget curl gcc g++
RUN make
RUN apk del make git wget curl gcc g++

ENTRYPOINT /go/src/github.com/reivaj05/micro-gen/scripts/start.sh
