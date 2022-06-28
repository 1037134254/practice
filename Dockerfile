version: "3"
# docker run -d -p 63xx:xxxx --network -v xxx:xxx
services:
 redis:
 # 镜像
 image: redis
 # 端口
 ports:
 - "6379:6379"
 volumes:
 - /app/redis/redis.conf:/etc/redis/redis.conf
 - /app/redis/data:/data
 networks:
 - data_net
 command: redis-server /etc/redis/redis.conf

 mysql:
 image: mysql
 # 环境
 environment:
 MYSQL_ROOT_PASSWORD: 'root'
 MYSQL_ALLOW_EMPTY_PASSWORD: 'no'
 MYSQL_DATABASE: 'OnlinePractice'
 MYSQL_USER: 'root'
 MYSQL_PASSWORD: 'root1234'
 ports:
 - "3306:3306"
 volumes:
 - /app/mysql/db:/var/lib/mysql
 - /app/mysql/conf/my.cnf:/etc/my.cnf
 - /app/mysql/init:/docker-entrypoint-initdb.d
 networks:
 - data_net
 command: --default-authentication-plugin=mysql_native_password #解决外部无法访问

# 网络配置
networks:
 data_net: