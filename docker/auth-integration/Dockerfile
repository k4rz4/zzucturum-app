#syntax=docker/dockerfile:1
FROM golang:1.17 AS build-env

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# resolve dependency before copying whole source code
COPY ./go.mod .
COPY ./go.sum .

# copy other sources & build
COPY ./ .
WORKDIR /go/src/app/cmd/auth-integration
RUN go install
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/app
RUN ls /go/bin

FROM alpine:3.11 AS runtime-env
COPY --from=build-env /go/bin/app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]