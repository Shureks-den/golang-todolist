FROM golang:1.17 AS build

ADD . /todoServer
WORKDIR /todoServer
RUN go build ./cmd/main.go

FROM ubuntu:20.04

RUN apt-get -y update &&\
    apt-get install -y tzdata

ENV POSTGRES_VERSION 12

RUN apt-get -y update &&\
    apt-get install -y postgresql-${POSTGRES_VERSION}

USER postgres

RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER todolist WITH SUPERUSER PASSWORD 'password';" &&\
    createdb -O postgres todolist &&\
    /etc/init.d/postgresql stop

RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/${POSTGRES_VERSION}/main/pg_hba.conf
RUN echo "listen_addresses='*'" >> /etc/postgresql/${POSTGRES_VERSION}/main/postgresql.conf

EXPOSE 5432

VOLUME ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

USER root

WORKDIR /usr/src/todoServer

COPY . .

COPY --from=build /todoServer/main .

ENV PGPASSWORD password

CMD service postgresql start &&\
    psql -h localhost -d todolist -U todolist -p 5432 -a -q -f ./database/db.sql &&\
    ./main