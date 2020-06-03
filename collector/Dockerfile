FROM golang:latest
MAINTAINER clianor <ejrgus94160@gmail.com>

RUN apt-get update
RUN apt-get install vim -y

WORKDIR /
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /
RUN chmod +x wait-for-it.sh

WORKDIR /go
ADD . .
RUN chmod +x start.sh
