package repository

import (
    "context"
    "os"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var uri = os.Getenv("MONGO_DB_CONNECTION")


func GetEvents() {
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(context.TODO(), opts)

    if err != nil {
        panic(err)
    }
    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
            panic(err)
        }
    }()

    var result bson.M
    if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
        panic(err)
    }

}


