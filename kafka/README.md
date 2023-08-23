# Kafka

## [环境配置](https://developer.confluent.io/quickstart/kafka-docker/)
创建 docker-compose.yml，在该文件中需要创建 zookeeper 和 kafka container
```yml
version: '3'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:7.3.2
    container_name: zookeeper
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000

  broker:
    image: confluentinc/cp-kafka:7.3.2
    container_name: broker
    ports:
    # To learn about configuring Kafka for access across networks see
    # https://www.confluent.io/blog/kafka-client-cannot-connect-to-broker-on-aws-on-docker-etc/
      - "9092:9092"
    depends_on:
      - zookeeper
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: 'zookeeper:2181'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_INTERNAL:PLAINTEXT
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092,PLAINTEXT_INTERNAL://broker:29092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1
```
kafka 默认是本地访问，倘若需要外部访问，将 KAFKA_ADVERTISED_LISTENERS 中的地址修改为公网地址
```yml
  KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://172.16.60.77:9092,PLAINTEXT_INTERNAL://broker:29092
```

使用 docker-compose 启动 kafka 和 zookeeper
```shell
docker-compose up -d
```

Linux 防火墙放开 9092 和 29092 端口
```shell
firewall-cmd --zone=public --add-port=9092/tcp --permanent
firewall-cmd --zone=public --add-port=29092/tcp --permanent
systemctl restart firewalld
```

## docker-compose 常用命令
1. docker-compose 运行容器：`docker-compose up -d`
2. docker-compose 关闭容器：`docker-compose down`
3. 查看日志：`docker-compose logs`
4. 查看指定容器日志：`docker-compose logs container-name`
5. 查看端口：`netstat -anp | grep port`
6. 重启 docker： `systemctl restart  docker`
7. 重启防火墙：`systemctl start firewalld`

## Golang Kafka client
1. https://github.com/segmentio/kafka-go

