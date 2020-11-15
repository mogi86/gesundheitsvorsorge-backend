FROM golang:1.15

WORKDIR /go/src/app
COPY . .

RUN make build

CMD ["./server"]