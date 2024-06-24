FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct


RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o user-rpcserver ./apps/user/rpcserver/user.go

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
RUN chmod +x user-rpcserver

# 设置容器启动时执行的命令
ENTRYPOINT ["./user-rpc", "-f", "./etc/user.yaml"]







































#FROM alpine:3.18
#
## 添加时区处理
#RUN echo -e "https://mirrors.aliyun.com/alpine/v3.15/main\nhttps://mirrors.aliyun.com/alpine/v3.15/community" > /etc/apk/repositories && \
#    apk update &&\
#    apk --no-cache add tzdata && \
#    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
#    echo "Asia/Shanghai" >  /etc/timezone
#ENV TZ=Asia/Shanghai
#
#ARG SERVER_NAME=user
#ARG SERVER_TYPE=rpcserver
#
#ENV RUN_BIN bin/${SERVER_NAME}-${SERVER_TYPE}
#ENV RUN_CONF /${SERVER_NAME}/conf/${SERVER_NAME}.yaml
#
## 这个关键词的意思是复制的意思，可以将宿主机中的内容复制到容器中
## 命令 左边是宿主机的目录，右边是容器目录
#RUN mkdir /$SERVER_NAME && mkdir /$SERVER_NAME/bin && mkdir /$SERVER_NAME/conf
#
## 复制编译后的二进制文件
#COPY ./bin/$SERVER_NAME-$SERVER_TYPE /$SERVER_NAME/bin/
#
## 复制配置文件
#COPY ./apps/$SERVER_NAME/$SERVER_TYPE/etc/dev/$SERVER_NAME.yaml /$SERVER_NAME/conf/
#
## 为二进制提供执行权限
#RUN chmod +x /$SERVER_NAME/bin/$SERVER_NAME-$SERVER_TYPE
#
## 该命令指定容器会默认进入那个目录，如我们每次进入服务器的时候会自动进入root目录一样的作用
#WORKDIR /$SERVER_NAME
#
## 这个命令可以让我们的docker容器在启动的时候就执行下面的命令
## 与CMD不同之处是，在docker run 后跟的命令不能替换它，它仍然会启动的时候执行
## ENTRYPOINT ["$RUN_BIN", "-f", "$RUN_CONF"] // 这种写法不支持对环境变量的解析，
##您正在使用ENTRYPOINT 的exec形式。与shell表单不同，exec表单不会调用命令shell。这意味着正常的外壳处理不会发生。例如，ENTRYPOINT [ "echo", "$HOME" ]
## 将不会在$ HOME上进行变量替换。如果要进行shell处理，则可以使用shell形式或直接执行shell，例如：ENTRYPOINT [ "sh", "-c", "echo $HOME" ]。
##当使用exec表单并直接执行shell时（例如在shell表单中），是由shell进行环境变量扩展，而不是docker。（来自Dockerfile参考）
##
#
#ENTRYPOINT $RUN_BIN -f $RUN_CONF
