FROM golang:latest

WORKDIR $GOPATH/src/github.com/Yingzhixin/gin_demo
RUN git clone https://github.com/Yingzhixin/gin_demo.git
ENV GOPROXY https://goproxy.io
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./gin_demo"]

# remember to change the registry-mirrors
#
#
# cd /etc/docker
# vim daemon.json
#
#
# {
#     "registry-mirrors":["https://registry.docker-cn.com"]
# }
