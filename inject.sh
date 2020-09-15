#!/usr/bin/env bash

kubectl get deployments | grep nginx  | awk '{print $1}' | xargs -I {} kubectl patch deployment {} --patch "$(cat latency-setting-sidecar-single.yaml)"
