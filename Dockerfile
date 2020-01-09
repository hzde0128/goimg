FROM golang:1.13 as builder

ADD . /go/src/github.com/hzde0128/goimg

WORKDIR /go/src/github.com/hzde0128/goimg

RUN go build

FROM debian:buster

COPY --from=builder /go/src/github.com/hzde0128/goimg/goimg /app


EXPOSE 8080


VOLUME /data

CMD ["/app"]
