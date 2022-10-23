FROM golang:1.19.2-alpine3.16

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && apk add tzdata bind-tools --no-cache
RUN apk add build-base
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod  vendor -v

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN GO15VENDOREXPERIMENT="1" CGO_ENABLED="1" GOOS=linux GOARCH=amd64 go build -trimpath -o /usr/local/bin/client client/main.go
RUN GO15VENDOREXPERIMENT="1" CGO_ENABLED="1" GOOS=linux GOARCH=amd64 go build -trimpath -o /usr/local/bin/server main.go
ADD conf/app.conf  /usr/local/bin/conf/app.conf
WORKDIR /usr/local/bin
