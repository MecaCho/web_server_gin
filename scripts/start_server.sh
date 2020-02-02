#!/usr/bin/env bash

export DBCON="root:xxxxxx@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Asia%2FShanghai&parseTime=true"
export IP=0.0.0.0
export PORT=80
export LOG_FILE=demo.log

/usr/bin/demo --alsologtostderr=true \
--serverIp=${IP} \
--serverPort=${PORT} \
--dbConnect=${DBCON} \
--debug=false \
>> ${LOG_FILE} 2>&1 &
