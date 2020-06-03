FROM arm64v8/ubuntu:18.04

ENV MACHBASE_HOME=/machbase_home
ENV PATH=$MACHBASE_HOME/bin:$PATH 
ENV LD_LIBRARY_PATH=$MACHBASE_HOME/lib:$LD_LIBRARAY_PATH

# 의존성 라이브러리 설치 및 업데이트
RUN apt-get update -y && apt-get upgrade -y \
	&& apt-get install wget -y \
	&& apt-get install --reinstall tzdata -y \
	&& apt-get install libssl1.0.0 -y

# Time Zone 설정
RUN cp -p /usr/share/zoneinfo/Asia/Seoul /etc/localtime

WORKDIR /machbase_home
RUN wget http://dl.machbase.com/dist/machbase-edge-5.5.7.official-LINUX-ARM_CORTEX_A53-64-release.tgz
RUN tar zxf machbase-edge-5.5.7.official-LINUX-ARM_CORTEX_A53-64-release.tgz && rm -f machbase-edge-5.5.7.official-LINUX-ARM_CORTEX_A53-64-release.tgz
RUN machadmin -c

# RUN echo machadmin -u > start.sh && echo bash >> start.sh
RUN echo machadmin -u > start.sh && echo /bin/bash >> start.sh
# ENTRYPOINT sh start.sh
ENTRYPOINT ["sh", "start.sh"]
CMD ["/bin/bash"]

EXPOSE 5656
