FROM golang
#RUN apt-get update && apt-get install -y tzdata
RUN mkdir -p /app/logs
RUN ls -lka /app
# 设置工作目录
WORKDIR /app
COPY ./ /app
ENV GO111MODULE on
ENV GOPROXY=https://goproxy.cn,direct
ENV LANG C.UTF-8
ENTRYPOINT ["/app/klaytn-adapter"]
