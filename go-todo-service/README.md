# go-mongo

Go with mongodb for beginners.

Golang is a popular programming language known for its efficiency and concurrency support. MongoDB is a widely-used NoSQL database. Integrating Go with MongoDB allows you to build efficient and scalable applications. Below, I'll outline the steps to connect Go with MongoDB and perform basic operations.

## Run
```bash
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson

go mod init github.com/sharantechuser/gotodoservice
go mod tidy
go build

```

## Get all todos
####  curl http://localhost:4000/api/todos

## Create Todo

#### http://localhost:4000/api/todo

```bash
{
  "user": "sharan", 
  "description": "description",
  "datetime": "2021-02-18T21:54:42.123Z"
}
```

## Update Todo

#### curl -i -X PUT http://localhost:4000/api/todo/651dce8ee500ba78ea40b59d

```bash

  {
    "_id": "651dce8ee500ba78ea40b59d",
    "user": "sharan", 
    "description": "description",
    "datetime": "2021-02-18T21:54:42.123Z"
}
```
## Delete Todo

#### curl -i -X DELETE http://localhost:4000/api/todo/651dce8ee500ba78ea40b59d


