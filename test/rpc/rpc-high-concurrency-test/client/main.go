package main

import (
	"go-tests/rpc/rpc-high-concurrency-test/client/mymodules/myconfig"
	"log"
	"net/rpc"
	"sync"
	"time"
)

// Params 参数结构体, 字段必须是导出
type Params struct {
	Width, Height int
}

func main() {
	var c myconfig.MyConf
	c.GetConf()

	log.Println("loop concurrent: ", c.Concurrent)
	log.Println("rpc ipport: ", c.IPPort)
	log.Println("tryTimes: ", c.TryTimes)
	for i := c.TryTimes; i > 0; i-- {
		log.Printf("starting rpc loop in %d seconds...\n", i)
		time.Sleep(time.Duration(time.Second))
	}

	var wg sync.WaitGroup
	wg.Add(2 * c.Concurrent)

	for i := 0; i < c.Concurrent; i++ {
		go func() {
			defer wg.Done()

			// 连接远程 rpc 服务
			// 这里使用 Dial, http 方式使用 DialHTTP, 其他代码都一样
			rpc, err := rpc.Dial("tcp", c.IPPort)
			if err != nil {
				log.Fatal(err)
			}
			ret := 0
			// 调用远程方法
			// 注意第三个参数是指针类型
			err = rpc.Call("Rect.Area", Params{50, 100}, &ret)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(ret)
		}()

		go func() {
			defer wg.Done()

			// 连接远程 rpc 服务
			// 这里使用 Dial, http 方式使用 DialHTTP, 其他代码都一样
			rpc, err := rpc.Dial("tcp", c.IPPort)
			if err != nil {
				log.Fatal(err)
			}
			ret := 0
			// 调用远程方法
			// 注意第三个参数是指针类型
			err = rpc.Call("Rect.Perimeter", Params{50, 100}, &ret)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(ret)
		}()
	}

	wg.Wait()
	log.Println("main out...")
}
