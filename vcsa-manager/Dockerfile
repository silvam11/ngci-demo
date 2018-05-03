FROM golang:1.9.1-alpine

EXPOSE 8080

RUN apk add --update git; \
    mkdir -p ${GOPATH}/vcsa-manager; \
    go get -u github.com/gorilla/mux

WORKDIR ${GOPATH}/vcsa-manager/

COPY vcsa-manager.go ${GOPATH}/vcsa-manager/

RUN go build -o vcsa-manager .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/vcsa-manager/vcsa-manager .
CMD [ "/app/vcsa-manager" ]

