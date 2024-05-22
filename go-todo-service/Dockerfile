FROM golang:1.22

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build -o gotodoservice .

EXPOSE 4000

ENV MONGO_DATABASE=gotodoservicetest
ENV MONGO_PORT=27017

CMD [ "./gotodoservice" ]