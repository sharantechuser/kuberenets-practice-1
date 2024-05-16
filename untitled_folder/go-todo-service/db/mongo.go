package db

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/sharantechuser/gotodoservice/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client  *mongo.Client
	db      *mongo.Database
	context context.Context
}

func NewMongo() dbHandle {
	ctx := context.Background()
	m := &Mongo{context: ctx}
	if !m.initDB() {
		m.Client = nil
	}
	return m
}

func (m *Mongo) initDB() bool {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(m.context, clientOptions)
	if err != nil {
		log.Println(err)
		return false
	}
	m.Client = client
	m.db = client.Database(config.MongoDatabase)
	return true
}

func (m *Mongo) Create(i interface{}) error {

	coll_name := reflect.TypeOf(i)
	collection := m.db.Collection(coll_name.Name())

	insertResult, err := collection.InsertOne(m.context, i)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Inserted document with ID:", insertResult.InsertedID)
	return err
}

func (m *Mongo) GetAll(coll_name string, results interface{}) (interface{}, error) {

	collection := m.getCollection(coll_name)

	// var results []interface{}
	cursor, err := collection.Find(m.context, bson.D{{}})
	if err != nil {
		return results, err
	}

	if err = cursor.All(m.context, &results); err != nil {
		panic(err)
	}

	return results, err
}

func (m *Mongo) Update(_id string, v interface{}) (int64, error) {
	coll_name := reflect.TypeOf(v)
	collection := m.db.Collection(coll_name.Name())

	var doc bson.D
	data, err := bson.Marshal(v)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	err = bson.Unmarshal(data, &doc)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	id, _ := primitive.ObjectIDFromHex(_id)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", doc}}

	result, err := collection.UpdateOne(m.context, filter, update)
	if err != nil {
		log.Fatal(err)
		return result.MatchedCount, err
	}
	fmt.Println("Inserted document with ID:", result.MatchedCount)
	return result.MatchedCount, err
}

func (m *Mongo) Delete(_id string, v interface{}) (int64, error) {
	coll_name := reflect.TypeOf(v)
	collection := m.db.Collection(coll_name.Name())

	var doc bson.D
	data, err := bson.Marshal(v)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	err = bson.Unmarshal(data, &doc)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	id, _ := primitive.ObjectIDFromHex(_id)
	filter := bson.D{{"_id", id}}

	result, err := collection.DeleteOne(m.context, filter)
	if err != nil {
		log.Fatal(err)
		return result.DeletedCount, err
	}
	fmt.Println("Deleted document with ID:", result.DeletedCount)
	return result.DeletedCount, err
}

func (m *Mongo) getCollection(coll_name string) *mongo.Collection {
	return m.db.Collection(coll_name)
}

func (m *Mongo) IsValid() bool {
	return m.db != nil
}
