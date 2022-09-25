#!/bin/bash

docker build -t nginx-rp .
docker run -p 4000:80 nginx-rp