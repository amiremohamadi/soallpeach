FROM golang:latest
RUN go get -u github.com/labstack/echo/...
RUN mkdir -p /go/src/countme
COPY main.go /go/src/countme
WORKDIR /go/src/countme
RUN go build .
EXPOSE 80
ENTRYPOINT ./countme