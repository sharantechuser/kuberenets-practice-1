# kuberenets-practice-1

Created simple ToDo application using microservice architecture.


## Microservices 

1.  angular-frontend
2.  python-user-service
3.  go-todo-service
4.  MySQL
5.  MongoDB
6.  Jobs
6.  Apache Kafka

### 1.  angular-frontend
- The simple frontend microservice implemented in angular.

### 2.  python-user-service

- python-user-service is a backend service implemented in python.

### 3.  go-todo-service

- go-todo-service is a backend service implemented in Golang.

### 4.  MySQL

- MySQL is a database service. user-serive will make use of this DB to store user detail data.

### 5.  MongoDB

- MongoDB is a database service. todo-serive will make use of this DB to store todo data.

### 6.  Job

- Its a kubernetes job run every scheduled time. 

### 7.  Apache Kafka

- Apache Kafka is a messaging broker. Job will publish kafka the message and go-todo-service consume the message.


# Implementation Progress
- [x] angular-frontend
- [x] python-user-service
- [x] go-todo-service
- [ ] MySQL
- [ ] MongoDB
- [ ] Jobs
- [ ] Apache Kafka