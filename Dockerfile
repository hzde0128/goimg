FROM golang:1.13 as builder
ADD . /go/src/github.com/hzde0128/goimg
WORKDIR /go/src/github.com/hzde0128/goimg
RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" -o app ./


FROM scratch
COPY --from=builder /go/src/github.com/hzde0128/goimg/app /app
EXPOSE 8080
VOLUME /data
CMD ["/app"]
