#
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
#
# docker build -t nsip/hitsproxy .
# docker run -d -p8089:8089 nsip/hitsproxy
############################
# STEP 1 build executable binary
############################
FROM golang:1.13-stretch as builder
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
#RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
# Create appuser
#ENV USER=appuser
#ENV UID=10001
## See https://stackoverflow.com/a/55757473/12429735RUN 
#RUN adduser \    
#    --disabled-password \    
#    --gecos "" \    
#    --home "/nonexistent" \    
#    --shell "/sbin/nologin" \    
#    --no-create-home \    
#    --uid "${UID}" \    
#    "${USER}"WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Fetch dependencies.
# Using go mod with go 1.11
RUN go get
RUN go get github.com/labstack/echo/middleware
RUN go get github.com/nsip/sifclient-go
RUN go get github.com/nsip/sifdata-go
#RUN go mod verify
# Build the binary
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hitsproxy cmd/hitsproxy/main.go
############################
# STEP 2 build a small image
############################
FROM debian:stretch
## Import from builder.
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /etc/group /etc/group
## Copy our static executable
COPY --from=builder /go/bin/hitsproxy /go/bin/hitsproxy
# Use an unprivileged user.
# (TODO) USER appuser:appuser
# Run the hello binary.
CMD ["/go/bin/hitsproxy"]
