FROM golang:1.13 as builder

ADD . /go/src/goimg

WORKDIR /go/src/goimg

RUN go build

FROM debian:buster

COPY --from=builder /go/src/goimg/goimg /app


EXPOSE 8080


VOLUME /data

CMD ["/app"]
