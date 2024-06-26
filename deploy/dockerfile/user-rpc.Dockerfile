FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct


RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o user-rpc ./apps/user/rpc/user.go

FROM alpine:latest

# 正确配置 Alpine Linux 的软件源，并安装必要的软件包
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache ca-certificates tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo 'Asia/Shanghai' >/etc/timezone

# 设置工作目录
WORKDIR /app

# 从构建阶段名为 builder 的镜像中复制文件
COPY --from=builder /app/user-rpc ./
COPY --from=builder /app/apps/user/rpc/etc/user.yaml ./etc/user.yaml

# 设置可执行权限
RUN chmod +x user-rpc

# 设置容器启动时执行的命令
ENTRYPOINT ["./user-rpc", "-f", "./etc/user.yaml"]