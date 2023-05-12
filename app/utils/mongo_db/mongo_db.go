package mongo_db

import (
	"context"
	"fmt"

	"component-combine/app/model/metadata"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo 连接到 MongoDB
func ConnectMongo(envInfo *env.EnvConfig, collectionName string) (*mongo.Collection, *mongo.Client) {
	// Set client options 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://root:c75b4796f6c6e00c@47.100.184.170:27017/?authSource=admin&readPreference=primary&directConnection=true")
	// Connect to MongoDB 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Check the connection 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	// 指定获取要操作的数据集
	//collection := client.Database(envInfo.Mongo.Database).Collection(envInfo.Mongo.Collection)
	collection := client.Database(envInfo.Mongo.Database).Collection(collectionName)

	return collection, client
}

// DisconnectMongo 断开连接
func DisconnectMongo(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

// InsertOneMongo 插入单条数据
func InsertOneMongo(collection *mongo.Collection, character interface{}) {
	insertResult, err := collection.InsertOne(context.TODO(), character)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// FindOneMongo 查找一个文档
func FindOneMongo(collection *mongo.Collection, result interface{}) {

	filter := bson.D{{"metadata.attributes.value", "s-background"}}
	//filter := bson.D{{"created_at", "2022-11-07T15:44:14.774Z"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}

// FindManyMongo 查找多个文档
func FindManyMongo(collection *mongo.Collection, filter bson.D) interface{} {

	// 要查找多个文档，使用 Find()
	findOptions := options.Find()
	findOptions.SetLimit(1000)

	// 定义一个切片用来存储查询结果
	var results []metadata.CharacterDetail
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		panic(err)
	}

	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem metadata.CharacterDetail
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		panic(err)
	}

	// 完成后关闭游标
	cur.Close(context.TODO())

	// 修改更新时间
	//updateAtAtStr := generator.CreateISOTime(generator.UpdateAt)

	//fmt.Printf("Found multiple documents (array of pointers): %#v\n", results)
	return results

}

func QueryFilter(query []string) bson.D {
	//var AndQuery []map[string]interf{}
	//for i := 0; i < len(body.Properties); i++ {
	//	log.Println(body.Properties[i])
	//	currentCondition := bson.M{"properties": bson.M{"$elemMatch": bson.M{"propertyName": body.Properties[i].PropertyName, "propertyValue": body.Properties[i].PropertyValue}}}
	//	AndQuery = append(AndQuery, currentCondition)
	//}
	filter := bson.D{
		//{"metadata.attributes.value", "cowboy"},
		//{"metadata.attributes.value", "earphone"},
	}
	for i := 0; i < len(query); i++ {
		//f1 :=
		filter = append(filter, bson.E{Key: "metadata.attributes.value", Value: query[i]})
	}

	return filter
}
