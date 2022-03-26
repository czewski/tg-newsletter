package mongoCon

import (
	"context"
	"log"
	"time"

	"github.com/czewski/tg-newsletter/pkg/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (client *mongo.Client) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	//Connect to MongoDb
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://go:" + constants.ReadKey("mongoPwd") + "@cluster0.gs1ad.mongodb.net/myFirstDatabase?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//Handle disconnect
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	return client
}

func InsertInCollection(client *mongo.Client) (id interface{}) {
	//Obtain collection - later need to change to obtain a certain user, like fila_12309132
	collection := client.Database("filas").Collection("fila1")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		log.Fatal(err)
	}
	id = res.InsertedID

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return id
}
