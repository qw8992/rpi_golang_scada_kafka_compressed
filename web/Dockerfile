FROM python:3.6
MAINTAINER clianor <ejrgus94160@gmail.com>

ENV PYTHONUNBUFFERED 1

RUN mkdir /code
WORKDIR /code

COPY requirements.txt /code/
RUN pip3 install -r requirements.txt
COPY . /code/

WORKDIR /
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /
RUN chmod +x wait-for-it.sh

WORKDIR /code
ADD . .
RUN chmod +x start.sh
