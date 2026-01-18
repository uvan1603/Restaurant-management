package database

import (
	"os"
	"time"
	"log"
	"fmt"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	Mongo_DB := os.Getenv("MONGODB_URL")

	client , err := mongo.NewClient(options.Client().ApplyURI(Mongo_DB))

	if err != nil{
		log.Fatal(err)
	}

	ctx , cancel := context.WithTimeout(context.Background() , 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to mongo db")

	return client
}


var Client *mongo.Client = DBinstance()

func OpenCollection (client *mongo.Client , collectionName string) *mongo.Collection {
	dbName := os.Getenv("MONGO_DB_NAME")
	var collection *mongo.Collection = client.Database(dbName).Collection(collectionName)
	return collection
}