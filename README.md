# kuberenets-practice-1

Created simple three-tier ToDo application using microservice architecture.


# Deployement 

In three ways of application deployment has been done

## 1.  Deployed in AWS  using docker
   #### angular-frontend  
    Docker build . -t angularfrontent

    Docker run -p 8000:4200 angularfrontent
   ####  python-user-service 
    Docker build . -t  pythonuserservice

    Docker run --name python-backend -p 5000:5000 --networ python-mysql-network -e MYSQL_ROOT_PASSWORD=admin -e MYSQL_HOST=mysql pythonuserservice

    Docker run --name mysql -p 3306:3306 --networ python-mysql-network -e MYSQL_ROOT_PASSWORD=admin mysql:8.2
## 2.  Deployed in AWS using docker-compose
   #### angular-frontend  
    Docker build . -t angularfrontent

    Docker run -p 5000:5000 angularfrontent

   ####  python-user-service 
    docker-compose up --build
## 3.  Deployed in AWS kubernetes
   #### angular-frontend  
    kubectl apply -f angular-frontend.yml
    kubectl apply -f angular-frontend-svc.yml

   ####  python-user-service 
    kubectl apply -f python-user-service.yml
    kubectl apply -f python-user-service-svc.yml
    kubectl apply -f pv.yml
    kubectl apply -f pvc.yml




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