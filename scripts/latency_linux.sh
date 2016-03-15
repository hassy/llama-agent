#!/bin/bash

sudo tc qdisc add dev eth0 root netem delay $1
