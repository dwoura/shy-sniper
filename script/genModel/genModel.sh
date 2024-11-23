#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh 数据库名 表名称
# 比如：
# ./genModel.sh lottery lottery
# 再将./genModel下的生成的文件剪切到对应服务的model目录中即可

#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=./genModel

# 数据库配置
host=127.0.0.1  
port=3306
dbname=$1
username=root
passwd=xxxxx

echo "开始创建库：$dbname 的表：$2"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}" -dir="${modeldir}" -cache=true --home="${template}" --style=goZero
