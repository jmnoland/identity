package repository

import (
    "context"
    "os"
    "time"

    "github.com/jmnoland/identity/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var uri = os.Getenv("MONGO_DB_CONNECTION")
var database = os.Getenv("MONGO_DB_NAME")
var serverAPI = options.ServerAPI(options.ServerAPIVersion1)
var opts = options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).SetBSONOptions(&options.BSONOptions{ DefaultDocumentM: true })

func GetEvents() ([]model.Event) {
    client, err := mongo.Connect(context.TODO(), opts)

    if err != nil {
        panic(err)
    }
    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
            panic(err)
        }
    }()

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    events := client.Database(database).Collection("events")
    cur, err := events.Find(context.TODO(), bson.D{})
    if err != nil {
        panic(err)
    }
    defer cur.Close(ctx)
    
    var eventList []model.Event = make([]model.Event, 0)
    for cur.Next(ctx) {
        var result model.Event
        err := cur.Decode(&result)
        if err != nil {
            panic(err)
        }
        eventList = append(eventList, result)

    }
    return eventList
}

func AddEvent(newEvent model.Event) {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    client, err := mongo.Connect(ctx, opts)
    if err != nil {
        panic(err)
    }
    defer func() {
        if err = client.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()
    defer cancel()

    events := client.Database(database).Collection("events")
    _, err = events.InsertOne(ctx, newEvent)
    if err != nil {
        panic(err)
    }
}

