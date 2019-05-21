
####===========进入到容器内备份数据库数据
docker exec -it edgex-club-mongo /bin/sh

mongoexport --db edgex-club --collection article --out article.json

mongoexport --db edgex-club --collection comment --out comment.json

mongoexport --db edgex-club --collection message --out message.json

mongoexport --db edgex-club --collection reply --out reply.json

mongoexport --db edgex-club --collection session --out session.json

mongoexport --db edgex-club --collection user --out user.json

###=============退出容器后，在宿主机中运行
docker cp edgex-club-mongo:/mongodb-data mongodb-data



###============导入数据
mongoimport --db edgex-club --collection article --authenticationDatabase edgex-club --username edgex-club-db --password Yanhua@1989 --file article.json


mongoimport --db edgex-club --collection comment --authenticationDatabase edgex-club --username edgex-club-db --password Yanhua@1989 --file comment.json


mongoimport --db edgex-club --collection message --authenticationDatabase edgex-club --username edgex-club-db --password Yanhua@1989 --file message.json


mongoimport --db edgex-club --collection reply --authenticationDatabase edgex-club --username edgex-club-db --password Yanhua@1989 --file reply.json


mongoimport --db edgex-club --collection session --authenticationDatabase edgex-club --username edgex-club-db --password Yanhua@1989 --file session.json


mongoimport --db edgex-club --collection user --authenticationDatabase edgex-club --username edgex-club-db --password Yanhua@1989 --file user.json
