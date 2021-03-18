FROM golang:1.14 as builder

RUN mkdir -p $GOPATH/src/github.com/Yangiboev/golang-mongodb-kafka
WORKDIR $GOPATH/src/github.com/Yangiboev/golang-mongodb-kafka

COPY . ./

RUN export CGO_ENABLED=1 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/mongo_consumer /


ENTRYPOINT ["/mongo_consumer"]
