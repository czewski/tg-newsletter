package mongoCon

import (
	"context"
	"fmt"
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
		fmt.Println(err)
	}

	return client
}

//InsertInCollection - Insert in user queue an element
func InsertInCollection(userID string, filaNew string, client *mongo.Client) (id interface{}) {
	//Obtain collection - later need to change to obtain a certain user, like fila_12309132
	collection := client.Database("filas").Collection("fila1")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$addToSet": bson.M{"news": bson.M{"$each": []string{filaNew}}}}

	//bson.M{{"user", userID}, {"news", filaNew}}

	//Still need to create a json for things to
	_, err := collection.UpdateOne(ctx, bson.M{"user": userID}, update)

	//res, err := collection.InsertOne(ctx, bson.M{"user": userID}, update)
	if err != nil {
		fmt.Println(err)
	}
	//id = res.InsertedID
	return id
}

//ObtainResult - Busca file do usuario (userID)
func ObtainResult(userID string, client *mongo.Client) []string {
	//Obtain collection - later need to change to obtain a certain user, like fila_12309132
	collection := client.Database("filas").Collection("fila1")

	var result struct {
		User string   `bson:"user"`
		Fila []string `bson:"news"`
	}

	//Later need to change this to filter for id
	filter := bson.D{{"user", userID}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Later need to change &result to a Struct in the format i need
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		fmt.Println(err)
	}
	return result.Fila
}

//InsertSentNewsToUser - Insert in user what news were sent
func InsertSentNewsToUser(url string, userID string, client *mongo.Client) (err error) {
	//Obtain collection - later need to change to obtain a certain user, like fila_12309132
	collection := client.Database("filas").Collection("sent")

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	update := bson.M{"$addToSet": bson.M{"sent": bson.M{"$each": []string{url}}}}

	//Still need to create a json for things to
	_, err = collection.UpdateOne(ctx, bson.M{"user": userID}, update)

	//res, err := collection.InsertOne(ctx, bson.M{"user": userID}, update)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

//CheckSentNews - Busca a fila de enviados para um usuario
func CheckSentNews(userID string, client *mongo.Client) (sent []string, err error) {
	collection := client.Database("filas").Collection("sent")

	var result struct {
		User string   `bson:"user"`
		Fila []string `bson:"sent"`
	}

	//{sent:{$all:["url"]}, user:"237725036"}
	//Later need to change this to filter for id
	//filter := bson.M{"$all": bson.M{"sent": url}, bson.M{"user": userID}}
	filter := bson.D{{"user", userID}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Later need to change &result to a Struct in the format i need
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		return sent, err
	}
	return result.Fila, err
}
