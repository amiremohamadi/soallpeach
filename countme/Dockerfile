FROM golang:latest as build
RUN go get -u github.com/labstack/echo/...
RUN mkdir -p /go/src/countme
COPY main.go /go/src/countme
WORKDIR /go/src/countme
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o countme main.go

FROM alpine
COPY --from=build /go/src/countme/countme .
EXPOSE 80
ENTRYPOINT ./countme