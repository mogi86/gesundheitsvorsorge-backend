FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN ls -la
RUN make build

CMD ["./server"]