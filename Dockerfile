FROM golang:1.14.4-alpine3.12 AS build_base

ARG HTTPS_PROXY
ARG HTTP_PROXY

#RUN apk add --no-cache gcc g++ make bash git
#RUN apk add --update nodejs yarn

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /
COPY go.mod .
COPY go.sum .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
#RUN echo http://mirrors.aliyun.com/alpine/v3.7/main/ >> /etc/apk/repositories
RUN apk update

#RUN curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -
#RUN cat <<EOF >/etc/apt/sources.list.d/kubernetes.list \
#    deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main \
#    EOF
COPY config config

RUN apk add --no-cache gcc g++ make bash git
RUN go mod download
COPY main.go /main.go
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o main

##todo: multistage-build
#FROM golang:1.14.4-alpine3.12
#COPY main /main
#WORKDIR /
#CMD ["/bin/bash"]
#ENTRYPOINT ["sh", "-c", "sleep 5m"]
RUN apk add iproute2
EXPOSE 8092

ENTRYPOINT ["./main"]
