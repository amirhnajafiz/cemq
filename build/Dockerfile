# from golang alpine image
FROM golang:1.22-alpine as builder

WORKDIR /src/app

# copy cmd, internal, and pkg
COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/
COPY main.go main.go

# copy dependencies
COPY go.mod go.sum ./

# compile go program
RUN go build -o /cemq

# second stage
FROM alpine:latest

WORKDIR /usr/bin/cemq

COPY --from=builder /cemq /cemq

ENTRYPOINT [ "./cemq" ]
