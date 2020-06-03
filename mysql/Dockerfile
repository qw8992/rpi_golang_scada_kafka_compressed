FROM arm64v8/alpine:latest

ENV MYSQL_VERSION=10.3.17-r0

RUN apk update && apk upgrade && apk add --no-cache tzdata && cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && echo "Asia/Seoul" > /etc/timezone && apk del tzdata

RUN apk add --no-cache mysql=${MYSQL_VERSION} mysql-client=${MYSQL_VERSION} mariadb-server-utils=${MYSQL_VERSION} && \
	addgroup mysql mysql && \
	mkdir /scripts && \
	rm -rf /var/cache/apk/*

COPY ./startup.sh /scripts/startup.sh
RUN chmod +x /scripts/startup.sh

VOLUME ["/var/lib/mysql"]
EXPOSE 3306

ENTRYPOINT ["/scripts/startup.sh"]
