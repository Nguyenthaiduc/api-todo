package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//mongodbInstance
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MongoIns MongoInstance

//connect Database
func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // cancel timeout

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect Successfully")
	fmt.Sprintln("Connect Successfully")

	MongoIns = MongoInstance{
		Client: client,
		Db : client.Database(os.Getenv("DATABASE_NAME")),
	}

}
