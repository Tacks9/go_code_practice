package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("【UDP SOCKET SERVICE】")

	service := "127.0.0.1:9001"
	// 获取一个UDPAddr
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	// 监听端口
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for{
		handlerUDPClient(conn)
	}


}

func handlerUDPClient(conn *net.UDPConn) {
	var buf [512]byte
	// 读取连接
	_,addr , err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	daytime := time.Now().String()
	// 写入当前时间
	conn.WriteToUDP([]byte(daytime), addr)
}

// 处理错误
func checkError(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
