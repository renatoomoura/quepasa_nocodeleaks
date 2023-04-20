FROM golang:1.19-stretch

WORKDIR /usr/src/app/
COPY . /usr/src/app/
RUN chmod 755 /usr/src/app/docker-entrypoint.sh
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT ["/usr/src/app/docker-entrypoint.sh"]
CMD ["sh"]