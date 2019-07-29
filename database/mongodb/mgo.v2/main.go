package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbServerIP string = "192.168.99.100"

// Student ...
type Student struct {
	Name   string `bson:"name"`
	Age    int    `bson:"age"`
	Sid    string `bson:"sid"`
	Status int    `bson:"status"`
}

// Per ...
type Per struct {
	Per []Student
}

func main() {
	succeeded := createAndInsertDB()
	if !succeeded {
		log.Println("createAndInsertDB failed")
		return
	}

	log.Println("createAndInsertDB succeeded")

	succeeded = findOne()
	if !succeeded {
		log.Println("findOne failed")
		return
	}

	log.Println("findOne succeeded")

	succeeded = findAll()
	if !succeeded {
		log.Println("findAll failed")
		return
	}

	log.Println("findAll succeeded")

	succeeded = updateData()
	if !succeeded {
		log.Println("updateData failed")
		return
	}

	log.Println("updateData succeeded")

	succeeded = delData()
	if !succeeded {
		log.Println("delData failed")
		return
	}

	log.Println("delData succeeded")
}

func delData() bool {
	mongo, err := mgo.Dial(dbServerIP) // 建立连接
	defer mongo.Close()
	if err != nil {
		log.Panic(err)
	}

	client := mongo.DB("mydb_tutorial").C("t_student")
	if client == nil {
		log.Println("Got client empty")
		return false
	}

	//只删除一条
	err = client.Remove(bson.M{"sid": "s20180907"})
	if err != nil {
		log.Panic(err)
	}

	return true
}

func updateData() bool {
	mongo, err := mgo.Dial(dbServerIP) // 建立连接
	defer mongo.Close()
	if err != nil {
		log.Panic(err)
	}

	client := mongo.DB("mydb_tutorial").C("t_student")
	if client == nil {
		log.Println("Got client empty")
		return false
	}

	//只更新一条
	err = client.Update(bson.M{"status": 1}, bson.M{"$set": bson.M{"age": 20}})
	// _, err = client.UpdateAll(bson.M{"status": 1}, bson.M{"$set": bson.M{"age": 20}})
	if err != nil {
		log.Panic(err)
	}

	return true
}

// 查找status为1的数据
func findAll() bool {
	mongo, err := mgo.Dial(dbServerIP) // 建立连接
	defer mongo.Close()
	if err != nil {
		log.Panic(err)
	}

	client := mongo.DB("mydb_tutorial").C("t_student")
	if client == nil {
		log.Println("Got client empty")
		return false
	}

	//每次最多输出15条数据
	iter := client.Find(bson.M{"status": 1}).Sort("_id").Skip(1).Limit(15).Iter()
	if iter == nil {
		log.Println("Got iter empty")
		return false
	}

	var stu Student
	var users Per
	for iter.Next(&stu) {
		users.Per = append(users.Per, stu)
	}

	if err := iter.Close(); err != nil {
		log.Panic(err)
	}
	fmt.Println(users)
	return true
}

func findOne() bool {
	mongo, err := mgo.Dial(dbServerIP) // 建立连接
	defer mongo.Close()
	if err != nil {
		log.Panic(err)
	}

	client := mongo.DB("mydb_tutorial").C("t_student")
	if client == nil {
		log.Println("Got client empty")
		return false
	}

	user := Student{}
	// 查找sid为 s20180907
	err = client.Find(bson.M{"sid": "s20180907"}).One(&user)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(user)
	return true
}

func createAndInsertDB() bool {
	mongo, err := mgo.Dial(dbServerIP) // 建立连接
	defer mongo.Close()
	if err != nil {
		log.Panic(err)
	}

	client := mongo.DB("mydb_tutorial").C("t_student") // 选择数据库和集合
	if client == nil {
		log.Println("Got client empty")
		return false
	}

	// 创建数据
	data := Student{
		Name:   "seeta",
		Age:    18,
		Sid:    "s20180907",
		Status: 1,
	}

	// 插入数据
	err = client.Insert(&data)
	if err != nil {
		log.Panic(err)
	}

	return true
}
