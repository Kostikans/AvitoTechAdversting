FROM golang:1.15 as avito_advertising
ENV GO111MODULE=on
WORKDIR /go/src/avito
COPY . /go/src/avito
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64   go build  ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=avito_advertising /go/src/avito /app
RUN chmod -R +x  .
EXPOSE 9000/tcp
ENTRYPOINT [ "/app/main" ]