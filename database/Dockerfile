FROM postgres:16-alpine3.19

# COPY . /root/
COPY ./script.sql /docker-entrypoint-initdb.d/

WORKDIR /root/