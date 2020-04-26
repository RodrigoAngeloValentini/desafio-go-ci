FROM golang:1.14.2-alpine as builder

WORKDIR /go/src/app
COPY ./src/sum .
RUN go build app.go

FROM scratch

COPY --from=builder /go/src/app .

CMD ["/app"]