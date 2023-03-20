FROM golang:latest As builder

# 设置环境变量
ENV GO111MODULE=on\
    CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64

# 移动到工作目录
WORKDIR /build

# 将项目中的go.mod和go.sum复制到容器并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器
COPY . . 

# 将代码编译成可执行文件app
RUN go build -o app-test . 

# 创建一个小镜像以传递
FROM debian:stretch-slim

# 拷贝基础配置设定
COPY wait-for.sh .
COPY config.yaml .

# 将可执行文件复制到小镜像中
COPY --from=builder /build/app-test /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# 声明服务器端口
EXPOSE 8088

# 需要运行的命令
# ENTRYPOINT [ "/app", "config.yaml" ]