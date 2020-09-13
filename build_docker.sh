#!/usr/bin/env bash
#todo check dir is success
dir=`pwd`
docker build -t latencyinjection:v1 $dir