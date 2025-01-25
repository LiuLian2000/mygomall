FROM golang:1.22.11-alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 更新Alpine的软件源为国内的站点（例如清华大学）  
 
# 更新软件包列表并安装Bash  
RUN apk update && apk upgrade && apk add --no-cache bash bash-doc bash-completion  
  
# 清理缓存以减小镜像体积  
RUN rm -rf /var/cache/apk/*  
  
# 设置默认的shell为Bash  
RUN echo "/bin/bash" >> /etc/shells   

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

RUN chmod +x ./build.sh

RUN ./build.sh

FROM debian:12-slim

# 从builder 拷贝到当前目录
COPY --from=builder /build/output /

RUN set -eux; \
    mv /etc/apt/sources.list /etc/apt/sources.list.bak; \
    echo "deb http://mirrors.aliyun.com/debian/ stretch main contrib non-free" > /etc/apt/sources.list; \
    echo "deb http://mirrors.aliyun.com/debian/ stretch-updates main contrib non-free" >> /etc/apt/sources.list; \
    echo "deb http://mirrors.aliyun.com/debian-security stretch/updates main contrib non-free" >> /etc/apt/sources.list; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

