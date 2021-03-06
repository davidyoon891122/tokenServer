package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../bodyStruct"
)

var dbMap map[string]string = map[string]string{
	"test":  "testDB",
	"token": "tokenDB",
}

var dbErrorCodeMap map[string]int = map[string]int{
	"Duplicated key":                      11000,
	"'$'is missing in the update process": 00001,
}

var merr mongo.WriteException

func connectMongo(dbname string, colname string) *mongo.Collection {
	clientOption := options.Client().ApplyURI("mongodb://localhost:32774/?connect=direct") //?replicaSet=rs0 is not working.
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	database := client.Database(dbname)
	collection := database.Collection(colname)

	return collection
}

func WriteData(data interface{}) int {
	collection := connectMongo(dbMap["test"], "userInfo")

	switch data.(type) {
	case *bodyStruct.Body:
		bsonData := bson.M{
			"token":  data.(*bodyStruct.Body).Token,
			"userID": data.(*bodyStruct.Body).UserID,
		}
		_, err := collection.InsertOne(context.TODO(), bsonData)

		if err != nil {
			merr = err.(mongo.WriteException)
			fmt.Printf("Number of errors : %d\n", len(merr.WriteErrors))
			errCode := merr.WriteErrors[0].Code
			fmt.Println("errCode of MongoDB : ", errCode)
			return errCode
		}
	}
	return 0
}

func ReadData(userID string) interface{} {
	collection := connectMongo(dbMap["test"], "userInfo")
	var err error
	var cursor *mongo.Cursor
	//read all data or specific data with userID
	if userID == "" {
		cursor, err = collection.Find(context.TODO(), bson.D{{}})
	} else {
		cursor, err = collection.Find(context.TODO(), bson.M{"userID": userID})
	}

	if err != nil {
		panic(err)
	}

	var resultArray []bson.M

	for cursor.Next(context.TODO()) {
		var result bson.M

		err := cursor.Decode(&result)

		if err != nil {
			panic(err)
		} else {
			resultArray = append(resultArray, result)
		}
	}

	defer cursor.Close(context.TODO())

	return resultArray
}

func UpdateData(data interface{}) int {
	collection := connectMongo(dbMap["test"], "userInfo")

	switch data.(type) {
	case *bodyStruct.Body:
		bsonData := bson.M{
			"$set": bson.M{
				"token":  data.(*bodyStruct.Body).Token,
				"userID": data.(*bodyStruct.Body).UserID,
			},
		}

		filter := bson.M{"userID": data.(*bodyStruct.Body).UserID}
		result, err := collection.UpdateOne(context.TODO(), filter, bsonData)

		fmt.Println("result update : ", result)
		if err != nil {
			fmt.Println(err)
			return 00001
		}
	}
	return 0
}
