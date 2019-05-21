#!/bin/sh
set -e
mongod --auth --dbpath=/mongo/db &
sleep 3
	set +e
mongo /mongo/initmongo.js
wait
