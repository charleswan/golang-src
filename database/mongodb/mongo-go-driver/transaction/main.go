package main

import (
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	connectString := "mongodb://127.0.0.1/test"
	dbURL, err := url.Parse(connectString)
	if err != nil {
		panic(err)
	}

	//认证参数设置, 否则连不上
	opts := &options.ClientOptions{}
	opts.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    "test",
		Username:      "test",
		Password:      "123456"})

	client, err := mongo.Connect(context.Background(), connectString, opts)
	if err != nil {
		panic(err)
	}

	db := client.Database(dbURL.Path[1:])

	ctx := context.Background()
	defer db.Client().Disconnect(ctx)

	col := db.Collection("test")

	//先在事务外写一条id为“111”的记录
	_, err = col.InsertOne(ctx, bson.M{"_id": "111", "name": "ddd", "age": 50})
	if err != nil {
		fmt.Println(err)
		return
	}

	//第一个事务：成功执行
	db.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction()
		if err != nil {
			fmt.Println(err)
			return err
		}

		//在事务内写一条id为“222”的记录
		_, err = col.InsertOne(sessionContext, bson.M{"_id": "222", "name": "ddd", "age": 50})
		if err != nil {
			fmt.Println(err)
			return err
		}

		//在事务内写一条id为“333”的记录
		_, err = col.InsertOne(sessionContext, bson.M{"_id": "333", "name": "ddd", "age": 50})
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			return err
		} else {
			sessionContext.CommitTransaction(sessionContext)
		}
		return nil
	})

	//第二个事务：执行失败, 事务没提交, 因最后插入了一条重复id "111",
	err = db.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			fmt.Println(err)
			return err
		}

		//在事务内写一条id为“222”的记录
		_, err = col.InsertOne(sessionContext, bson.M{"_id": "444", "name": "ddd", "age": 50})
		if err != nil {
			fmt.Println(err)
			return err
		}

		//写重复id
		_, err = col.InsertOne(sessionContext, bson.M{"_id": "111", "name": "ddd", "age": 50})
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			return err
		} else {
			sessionContext.CommitTransaction(sessionContext)
		}
		return nil
	})
}

// 最终数据只有 "111","222","333" 三条, 事务测试成功。
