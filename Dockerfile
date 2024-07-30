FROM golang:1.21-alpine as builder

WORKDIR /go/src
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags "-s -w" -trimpath -o /go/bin/app

FROM alpine:3.12

ENV TZ=Asia/Shanghai
RUN apk add --no-cache alpine-conf dumb-init curl && \
    setup-timezone -z ${TZ} && \
    apk del alpine-conf && \
    rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/app /app

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/app"]