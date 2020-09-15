#!/usr/bin/env bash

ifconfig -s | awk '{print $1}' | xargs -I {} tc qdisc del dev {} root netem
ifconfig -s | awk '{print $1}' | xargs -I {} tc qdisc add dev {} root netem delay ${1}ms