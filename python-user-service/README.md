# user-service python REST API  using Flask and MySQL

This is a simple backend user-service. It store user detail in mysql database.

This is a multi containerized application. Python and MySQl container are run together using docker-compose

# Dockerized the app

## Docker python-user-service

```
FROM python:3.9

WORKDIR /app

COPY ./dependencies.txt /app

RUN pip3 install -r dependencies.txt

COPY . .

EXPOSE 5000

CMD [ "python", "main.py" ]
```

## docker-compose.yml
```
version: '3'

services:
  pythonuserservice:
    container_name: pythonuserservice
    build: 
      context: .
    restart: always
    ports:
      - "8000:5000"
    volumes:
      - ./:/app
    networks:
      - mysql-python-app-network
    depends_on:
      - dockerdbhost

  dockerdbhost:
    container_name: dockerdbhost
    image: mysql:latest
    ports:
      - "3306:3306"
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: pwd@123
    volumes:
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql
      - ./mysql-data:/var/lib/mysql 
    networks:
      - mysql-python-app-network

networks:
  mysql-python-app-network:
    name: mysql-python-app-network

volumes:
  mysql-data:

```

## Docker build an run 

```
docker-compose up --build   
```

## Docker remove container
```
docker-compose down -v   
```

## Create User

#### curl -i -X POST http://localhost:8000/user/add
```
{
  "username": "sharan",
  "password": "password"
}
```

## Get User

#### curl http://localhost:8000/user/users
```
{
  "user_id": 1,
  "username": "sharan",
  "password": "password"
}
```