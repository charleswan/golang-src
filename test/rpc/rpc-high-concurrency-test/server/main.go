package main

import (
	"log"
	"net"
	"net/rpc"
)

// Params 参数结构体, 字段必须是导出
type Params struct {
	Width, Height int
}

// Rect ...
type Rect struct{}

// Area 面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

// Perimeter 周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rect := new(Rect)
	// 注册 rpc 服务
	rpc.Register(rect)
	// 获取 tcpaddr
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	chkError(err)
	// 监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	chkError(err2)
	// 死循环处理连接请求
	for {
		conn, err3 := tcplisten.Accept()
		if err3 != nil {
			continue
		}
		// 使用 goroutine 单独处理 rpc 连接请求
		go rpc.ServeConn(conn)
	}
}
