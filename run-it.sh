#!/bin/sh
#开发环境
docker run -it -d -p 80:8080 --name edgex-club-prod edgex-club-prod:0.2.0

#生产环境
docker run -it -d -p 80:8080 -p 443:443 --name edgex-club-prod edgex-club-prod:0.2.0
