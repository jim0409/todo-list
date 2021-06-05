FROM golang:1.16.3-alpine3.13 as builder

ENV GOPATH /go/
ENV GO_WORKDIR $GOPATH/src/todo-list
ENV GOFLAGS=-mod=vendor

# ARG GitTag=`git rev-parse --short=6 HEAD`
ARG gitTag
# claim workdir and move to workdir loc
WORKDIR $GO_WORKDIR
ADD . $GO_WORKDIR

RUN apk add gcc
RUN go build -ldflags "-X main.gitcommitnum=$gitTag"


FROM alpine:3.13.5
ENV GOPATH /go/
ENV GO_WORKDIR $GOPATH/src/todo-list
WORKDIR /app

# copy binary into container
COPY --from=builder $GO_WORKDIR/todo-list todo-list
RUN mkdir config
ADD ./config/app.dev.ini ./config
RUN mkdir logger

CMD ["./todo-list"]