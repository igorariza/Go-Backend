package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// import "go.mongodb.org/mongo-driver/mongo"
// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// client, err := mongo.Connect(ctx, options.Client().ApplyURI(
//    "mongodb+srv://iariza:<password>@socialnetowkcluster.ddi19.mongodb.net/GroupArPost?retryWrites=true&w=majority",
// ))
// if err != nil { log.Fatal(err) }
//"mongodb+srv://iariza:iariza@socialnetowkcluster.ddi19.mongodb.net/GroupArPost?retryWrites=true&w=majority"

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://iariza:iariza@socialnetowkcluster.ddi19.mongodb.net/GroupArPost?retryWrites=true&w=majority")

func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa a BD")
	return client
}

func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
