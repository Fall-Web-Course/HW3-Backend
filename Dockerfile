FROM golang:1.16.3-alpine3.13  as core

WORKDIR /go/src/
COPY ./go.mod ./go.sum ./

ARG GOPROXY=https://goproxy.io,direct

# install dependencies
RUN sh -c 'go mod download && go mod verify'

# copy codes
COPY . .

# compile project
FROM core as builder

# build project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/hw3_backend

# make final image from alpine base and compiled binary
FROM alpine:3.13
COPY --from=builder /go/bin/hw3_backend /usr/bin/hw3_backend

ENTRYPOINT ["hw3_backend"]