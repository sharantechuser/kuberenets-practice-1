# MySQL docker container

Rrun a MySQL container using the docker run command. You'll need to specify the MySQL Docker image, set environment variables for MySQL configuration (like root password, database name, etc.), and expose the required ports.

## Dockerfile
```
FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=pwd@123

COPY ./init.sql /docker-entrypoint-initdb.d/
```

## Docker build
```
docker build -t mysql .
```


## Docker run
```
docker run --name mysql -p 3306:3306 mysql
```