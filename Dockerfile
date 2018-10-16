FROM golang:1.11.1-alpine as builder
RUN apk add git
RUN mkdir -p /go/src/build
ADD . /go/src/build/
WORKDIR /go/src/build
ENV GOPATH /go/
RUN go get -v gopkg.in/yaml.v1
RUN go build -v -o main .
RUN dos2unix cvs.yml

FROM alpine
RUN adduser -S -D -H -h /app appuser
COPY --from=builder /go/src/build/main /app/
COPY --from=builder /go/src/build/cvs.yml /app/

USER appuser
WORKDIR /app
ENTRYPOINT /app/main