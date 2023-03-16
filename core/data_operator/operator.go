package data_operator

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var MongoUri string

func InsertMany(db, col string, docs []interface{}) error {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(MongoUri).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = client.Database(db).Collection(col).InsertMany(context.Background(), docs)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
