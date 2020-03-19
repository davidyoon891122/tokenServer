package db

import (
    "fmt"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "../bodyStruct"
)

var dbMap map[string]string = map[string]string {
    "test": "testDB",
    "token": "tokenDB",
}



func connectMongo(dbname string, colname string) *mongo.Collection {
    clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOption)

    if err != nil {
        panic(err)
    }

    err = client.Ping(context.TODO(), nil)

    database := client.Database(dbname)
    collection := database.Collection(colname)
    
    return collection

}


func WriteData(data interface{}) {
    collection := connectMongo(dbMap["test"], "userInfo")

    fmt.Println("collection test : ", collection )
    switch data.(type) {
        case bodyStruct.Body:
            fmt.Println(data.(bodyStruct.Body).Token)
    }
}



