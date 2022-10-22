FROM golang:1.18.1-alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && apk add tzdata bind-tools --no-cache
ADD bin/server /usr/local/bin/server
ADD bin/client /usr/local/bin/client
ADD app.conf  /usr/local/bin/conf/app.conf
WORKDIR /usr/local/bin