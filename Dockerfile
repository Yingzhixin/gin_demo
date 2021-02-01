FROM golang

# RUN git clone https://github.com/Yingzhixin/gin_demo.git
# remember to change the "host" in /config/application.yml to your docker_mysql_IPAddress

WORKDIR $GOPATH/src/github.com/Yingzhixin/gin_demo
COPY . $GOPATH/src/github.com/Yingzhixin/gin_demo

ENV GOPROXY https://goproxy.io
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./gin_demo"]

# remember to change the registry-mirrors

# cd /etc/docker
# vim daemon.json
#
# {
#     "registry-mirrors" : [
#     "https://registry.docker-cn.com",
#     "https://docker.mirrors.ustc.edu.cn",
#     "http://hub-mirror.c.163.com",
#     "https://cr.console.aliyun.com/"
#   ]
# }
#
# privacy docker mirrors: https://ngbz95d3.mirror.aliyuncs.com