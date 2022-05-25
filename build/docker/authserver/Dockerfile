FROM BASE_IMAGE
MAINTAINER chenbo <chenbo@ghgame.cn>

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

WORKDIR /opt/permission
COPY permission-authserver /opt/permission/bin/

ENTRYPOINT ["/opt/permission/bin/permission-authserver"]
