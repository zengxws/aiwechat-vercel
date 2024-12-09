# 使用官方的Go语言镜像作为构建阶段的基础镜像
FROM golang:1.21 AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制项目的所有源代码
COPY . .

# 编译项目
RUN go build -o main .

# 使用一个更小的基础镜像来运行编译后的二进制文件
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制编译后的二进制文件
COPY --from=builder /app/main .

# 设置容器启动时执行的命令
CMD ["./main"]