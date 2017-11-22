FROM golang:alpine


WORKDIR /go/src/github.com/reivaj05/micro-gen

ADD . /go/src/github.com/reivaj05/micro-gen

RUN chmod +x ./scripts/*.sh
RUN apk -U add make git bash wget curl gcc g++
RUN make
RUN apk del make git wget curl gcc g++

ENTRYPOINT ./scripts/start.sh
