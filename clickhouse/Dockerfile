FROM yandex/clickhouse-server:22.1.3.7
RUN apt update && apt install -y openjdk-17-jre-headless zookeeperd
COPY ./start-zookeeper.sh /start-zookeeper.sh
ENTRYPOINT ["/bin/sh", "-c", "/start-zookeeper.sh && /entrypoint.sh"]
