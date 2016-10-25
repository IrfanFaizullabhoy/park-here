FROM golang:1.7.1-alpine

COPY . $GOPATH/src/github.com/IrfanFaizullabhoy/nest

WORKDIR $GOPATH/src/github.com/IrfanFaizullabhoy/nest/server

RUN go build

ENTRYPOINT ["./server"]