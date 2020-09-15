#!/usr/bin/env bash
#todo check dir is success

mkdir build-docker
cd build-docker
git clone https://github.com/frankwyw7/k8s-addtional-latency-injection
cd k8s-addtional-latency-injection
cp /root/.kube/config config

dir=`pwd`
docker build -t latencyinjection:v1 $dir