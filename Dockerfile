FROM golang:1.11-alpine AS builder

MAINTAINER Zhang huaqiao <yhzhq1989@163.com>

RUN cp /etc/apk/repositories /etc/apk/repositories.bak
RUN echo "https://mirrors.ustc.edu.cn/alpine/v3.6/main" > /etc/apk/repositories
RUN echo "https://mirrors.ustc.edu.cn/alpine/v3.6/community" >> /etc/apk/repositories
RUN cat /etc/apk/repositories

WORKDIR /go/src/edgex-club/

ENV GOPROXY https://goproxy.io

RUN apk update && apk --no-cache add ca-certificates && apk add git make

COPY . .

RUN CGO_ENABLED=0 GO111MODULE=on go build -o edgex-club-linux

FROM alpine:3.6

RUN cp /etc/apk/repositories /etc/apk/repositories.bak
RUN echo "https://mirrors.ustc.edu.cn/alpine/v3.6/main" > /etc/apk/repositories
RUN echo "https://mirrors.ustc.edu.cn/alpine/v3.6/community" >> /etc/apk/repositories
RUN cat /etc/apk/repositories

RUN apk update && apk --no-cache add ca-certificates

WORKDIR /edgex-club/
COPY --from=builder /go/src/edgex-club/edgex-club-linux .
COPY --from=builder /go/src/edgex-club/env ./env/
COPY --from=builder /go/src/edgex-club/static ./static/

EXPOSE 8080

ENTRYPOINT ["./edgex-club-linux", "-confpath=env/prod.yaml"]
