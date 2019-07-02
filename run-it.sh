#!/bin/sh
#开发环境，-v 参数同步container容器和宿主机时间一致
docker run -it -d -v /etc/localtime:/etc/localtime -p 80:8080 --name edgex-club-prod edgex-club-prod:0.2.0

#生产环境 -v 参数同步container容器和宿主机时间一致
docker run -it -d -v /etc/localtiome:/etc/localtime -p 80:8080 -p 443:443 --name edgex-club-prod edgex-club-prod:0.2.0
