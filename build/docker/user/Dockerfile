FROM BASE_IMAGE
MAINTAINER cb<fastopencn@gmail.com>

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

WORKDIR /opt/permission
COPY user /opt/user/bin/

ENTRYPOINT ["/opt/user/bin/user"]
