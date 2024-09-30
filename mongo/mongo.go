package mongo

import (
	"context"
	"fmt"
	"github.com/Ryeom/cosmos/log"
	"github.com/Ryeom/cosmos/util"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"reflect"
	"time"
)

var MongoClient *mongo.Client

func InitialiseMongo() {
	//main mongo connection
	ip := viper.GetString("mongo.ip")
	port := viper.GetString("mongo.port")
	user := viper.GetString("mongo.user")
	pw := viper.GetString("mongo.pw")
	MongoClient = newClient(ip, port, user, pw)
}

func newClient(ip, port, user, pw string) *mongo.Client {
	if !util.IsPass(ip, port) {
		log.Logger.Error(ip + ":" + port + " 통신 불가.")
		return nil
	}
	address := "mongodb://" + ip + ":" + port
	return newMongoClient(address)
}

func newMongoClient(key string) *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	clientOptions := options.Client().ApplyURI(key).SetMaxPoolSize(3)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Logger.Error("Client Connection %s", err)

	}
	_ = client.Ping(ctx, nil)
	if err != nil {
		log.Logger.Error("Client Ping %s", err)

	}
	return client
}

func ToBsonD(i interface{}) bson.D {
	d := bson.D{}
	e := reflect.ValueOf(&i).Elem()
	fieldNum := e.NumField()
	for i := 0; i < fieldNum; i++ {
		v := e.Field(i)
		t := e.Type().Field(i)
		fmt.Printf("[Name: %s] Type: %s | Value: %v\n",
			t.Name, t.Type, v.Interface())
		d = append(d, bson.E{Key: t.Name, Value: v.Interface()})
	}
	return d
}

/**************************************************/

func SelectAll(client *mongo.Client, where map[string]string) map[string]string {
	result := map[string]string{}
	var l []bson.E
	for i, v := range where {
		l = append(l, bson.E{Key: i, Value: v})
	}
	//E의 배열이 D
	collection := client.Database("dcloudmongodb1").Collection("SizeMeta")
	cursor, err := collection.Find(context.TODO(), bson.D(l))
	if err != nil {
		log.Logger.Error("Find %s", err)
	}

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Logger.Error("All %s", err)
	}

	for _, v := range results {
		fmt.Println(v)
	}
	return result
}

func InsertOne(collection string, data bson.D) error {
	c := MongoClient.Database("cosmos").Collection(collection)
	insertResult, err := c.InsertOne(context.TODO(), data)
	fmt.Println(insertResult)
	return err
}
