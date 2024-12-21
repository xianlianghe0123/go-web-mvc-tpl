FROM alpine:3.21
LABEL author=xianlianghe0123

RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

WORKDIR /app

COPY ./config/ ./config/
COPY ./deploy/bin/server ./

EXPOSE 80
ENTRYPOINT [ "./server" ]

#docker build -t  1.1.1.1:5000/demo-api:v1 --build-arg APP_CONF=config/prod.yaml --build-arg  APP_RELATIVE_PATH=./cmd/server/...  .
#docker run -it --rm --entrypoint=ash 1.1.1.1:5000/demo-api:v1
