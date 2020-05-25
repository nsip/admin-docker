# BUILD: docker build -t nsip/example_golang .
# TEST: docker run -it -p3000:3000 nsip/example_golang .
# RUN: docker run -d -p3000:3000 nsip/example_golang

###########################
# STEP 0 Get them certificates
############################
FROM alpine:latest as certs
RUN apk --no-cache add ca-certificates

############################
# STEP 1 build executable binary
############################
FROM golang:1.13-stretch as builder
COPY . .
RUN go get github.com/labstack/echo/middleware
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/example ./main.go

############################
# STEP 2 build a small image
############################
FROM debian:stretch
COPY --from=builder /go/bin/example /go/bin/example
# NOTE - make sure it is the last build that still copies the files
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /go/bin
CMD ["/go/bin/example"]
