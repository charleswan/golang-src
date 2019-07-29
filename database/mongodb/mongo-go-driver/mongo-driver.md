<!-- TOC -->

- [1. mongodb 官方的 golang 驱动基础使用](#1-mongodb-官方的-golang-驱动基础使用)
    - [1.1. 导入](#11-导入)
    - [1.2. 链接 mongo 服务](#12-链接-mongo-服务)
    - [1.3. 判断服务是否可用](#13-判断服务是否可用)
    - [1.4. 选择数据库和集合](#14-选择数据库和集合)
    - [1.5. 删除这个集合](#15-删除这个集合)
    - [1.6. 插入一条数据](#16-插入一条数据)
    - [1.7. 批量插入数据](#17-批量插入数据)
    - [1.8. 查询单条数据](#18-查询单条数据)
    - [1.9. 查询单条数据后删除该数据](#19-查询单条数据后删除该数据)
    - [1.10. 询单条数据后修改该数据](#110-询单条数据后修改该数据)
    - [1.11. 查询单条数据后替换该数据 (以前的数据全部清空)](#111-查询单条数据后替换该数据-以前的数据全部清空)
    - [1.12. 一次查询多条数据 (查询 createtime>=3, 限制取 2 条, createtime 从大到小排序的数据)](#112-一次查询多条数据-查询-createtime3-限制取-2-条-createtime-从大到小排序的数据)
    - [1.13. 查询集合里面有多少数据](#113-查询集合里面有多少数据)
    - [1.14. 查询集合里面有多少数据 (查询 createtime>=3 的数据)](#114-查询集合里面有多少数据-查询-createtime3-的数据)
    - [1.15. 修改一条数据](#115-修改一条数据)
    - [1.16. 修改多条数据](#116-修改多条数据)
    - [1.17. 删除一条数据](#117-删除一条数据)
    - [1.18. 删除多条数据](#118-删除多条数据)
- [2. golang 操作 mongodb 的驱动 mongo-go-driver 的事务支持和访问控制](#2-golang-操作-mongodb-的驱动-mongo-go-driver-的事务支持和访问控制)
    - [2.1. mongo-go-driver 测试事务](#21-mongo-go-driver-测试事务)

<!-- /TOC -->

# 1. mongodb 官方的 golang 驱动基础使用

## 1.1. 导入

```Bash
go get -u -v go.mongodb.org/mongo-driver/mongo
```

## 1.2. 链接 mongo 服务

```Go
if client, err = mongo.Connect(getContext(), url); err != nil {
    checkErr(err)
}
```

## 1.3. 判断服务是否可用

```Go
if err = client.Ping(getContext(), readpref.Primary()); err != nil {
    checkErr(err)
}
```

## 1.4. 选择数据库和集合

```Go
collection = client.Database("testing_base").Collection("howie")
```

## 1.5. 删除这个集合

```Go
collection.Drop(getContext())
```

## 1.6. 插入一条数据

```Go
if insertOneRes, err = collection.InsertOne(getContext(), howieArray[0]); err != nil {
    checkErr(err)
}
fmt.Printf("InsertOne 插入的消息 ID:%v\n", insertOneRes.InsertedID)
```

## 1.7. 批量插入数据

```Go
if insertManyRes, err = collection.InsertMany(getContext(), howieArray); err != nil {
    checkErr(err)
}
fmt.Printf("InsertMany 插入的消息 ID:%v\n", insertManyRes.InsertedIDs)
```

## 1.8. 查询单条数据

```Go
if err = collection.FindOne(getContext(), bson.D{{"name", "howie_2"}, {"age", 11}}).Decode(&howie); err != nil {
    checkErr(err)
}
fmt.Printf("FindOne 查询到的数据:%v\n", howie)
```

## 1.9. 查询单条数据后删除该数据

```Go
if err = collection.FindOneAndDelete(getContext(), bson.D{{"name", "howie_3"}}).Decode(&howie); err != nil {
    checkErr(err)
}
fmt.Printf("FindOneAndDelete 查询到的数据:%v\n", howie)
```

## 1.10. 询单条数据后修改该数据

```Go
if err = collection.FindOneAndUpdate(getContext(), bson.D{{"name", "howie_4"}}, bson.M{"$set": bson.M{"name": "这条数据我需要修改了"}}).Decode(&howie); err != nil {
    checkErr(err)
}
fmt.Printf("FindOneAndUpdate 查询到的数据:%v\n", howie)
```

## 1.11. 查询单条数据后替换该数据 (以前的数据全部清空)

```Go
if err = collection.FindOneAndReplace(getContext(), bson.D{{"name", "howie_5"}}, bson.M{"hero": "这条数据我替换了"}).Decode(&howie); err != nil {
    checkErr(err)
}
fmt.Printf("FindOneAndReplace 查询到的数据:%v\n", howie)
```

## 1.12. 一次查询多条数据 (查询 createtime>=3, 限制取 2 条, createtime 从大到小排序的数据)

```Go
if cursor, err = collection.Find(getContext(), bson.M{"createtime": bson.M{"$gte": 2}}, options.Find().SetLimit(2), options.Find().SetSort(bson.M{"createtime": -1})); err != nil {
    checkErr(err)
}
if err = cursor.Err(); err != nil {
    checkErr(err)
}
defer cursor.Close(context.Background())
for cursor.Next(context.Background()) {
    if err = cursor.Decode(&howie); err != nil {
        checkErr(err)
    }
    howieArrayEmpty = append(howieArrayEmpty, howie)
}
fmt.Printf("Find 查询到的数据:%v\n", howieArrayEmpty)
```

## 1.13. 查询集合里面有多少数据

```Go
if size, err = collection.Count(getContext(), nil); err != nil {
    checkErr(err)
}
fmt.Printf("Count 里面有多少条数据:%d\n", size)
```

## 1.14. 查询集合里面有多少数据 (查询 createtime>=3 的数据)

```Go
if size, err = collection.Count(getContext(), bson.M{"createtime": bson.M{"$gte": 3}}); err != nil {
    checkErr(err)
}
fmt.Printf("Count 里面有多少条数据:%d\n", size)
```

## 1.15. 修改一条数据

```Go
if updateRes, err = collection.UpdateOne(getContext(), bson.M{"name": "howie_2"}, bson.M{"$set": bson.M{"name": "我要改了他的名字"}}); err != nil {
    checkErr(err)
}
fmt.Printf("UpdateOne 的数据:%d\n", updateRes)
```

## 1.16. 修改多条数据

```Go
if updateRes, err = collection.UpdateMany(getContext(), bson.M{"createtime": bson.M{"$gte": 3}}, bson.M{"$set": bson.M{"name": "我要批量改了他的名字"}}); err != nil {
    checkErr(err)
}
fmt.Printf("UpdateMany 的数据:%d\n", updateRes)
```

## 1.17. 删除一条数据

```Go
if delRes, err = collection.DeleteOne(getContext(), bson.M{"name": "howie_1"}); err != nil {
    checkErr(err)
}
fmt.Printf("DeleteOne 删除了多少条数据:%d\n", delRes.DeletedCount)
```

## 1.18. 删除多条数据

```Go
if delRes, err = collection.DeleteMany(getContext(), bson.M{"createtime": bson.M{"$gte": 7}}); err != nil {
    checkErr(err)
}
fmt.Printf("DeleteMany 删除了多少条数据:%d\n", delRes.DeletedCount)
```

# 2. golang 操作 mongodb 的驱动 mongo-go-driver 的事务支持和访问控制

[golang 操作 mongodb 的驱动 mongo-go-driver 的事务支持和访问控制](https://blog.csdn.net/sdghchj/article/details/85249392)

mongodb 要支持事务, 需要满足以下条件: 

- 4.0 以上版本; 
- 安装后时以 replication set(复本集) 模式启动; 
- storageEngine 存储引擎须是 wiredTiger (支持文档级别的锁), 4.0 以上版本已经默认是这个, [参考](https://docs.mongodb.com/manual/reference/configuration-options/#storage-options)

## 2.1. mongo-go-driver 测试事务

驱动源码里, 连接 server 过程内会先生成连接池, 然后返回有一个 client 对象, 通过 client 对象可以对 server 里的数据库集合进行读写。但是任何读写操作本身是不带 session 对象的, 所以在操作前会先生成一个默认的 session 对象, 然后再从连接池中取一个连接来进行通信。而事务相关的接口是在 session 接口内, 包括 Transaction 的 Start、Abort、Commit, 但 session 接口里并没有其它 CRUD 相关方法。

显然, 只有如下 closure 闭包方式才能使用事务接口
