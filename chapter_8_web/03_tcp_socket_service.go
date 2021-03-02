/*
 * @Descripttion: TCP Socket Service  可以使用telnet来尝试连接一下
 * @Author: tacks321@qq.com
 * @Date: 2021-03-02 16:04:03
 * @LastEditTime: 2021-03-02 16:48:00
 */
package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("【TCP SOCKET SERVICE】")

	service := "127.0.0.1:7777"
	// ResolveTCPAddr 获取一个 TCPAddr
	tcpAddr , err := net.ResolveTCPAddr("tcp4", service)
	/*
		//数据结构类型
		type TCPAddr struct {
			IP IP
			Port int
			Zone string // IPv6 scoped addressing zone
		}
	*/
	fmt.Println(tcpAddr)
	checkError(err)


	// 监听端口
	listener,err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// 循环接收请求
	for  {
		conn, err := listener.Accept()
		if err != nil {
			// 有错误的时候跳过，服务端记录信息
			continue
		}
		// 起协程进行处理 并发处理
		go handlerClient(conn)


	}

}

// 业务处理逻辑
func handlerClient(conn net.Conn) {
	// 处理超时的情况
	// conn.SetReadDeadline() 设置了超时，当一定时间内客户端无请求发送，conn 便会自动关闭，下面的 for 循环即会因为连接已关闭而跳出
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)

	// 关闭客户端连接
	defer conn.Close()

	for {
		// 不断接收客户端的请求
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		if readLen == 0 {
			break
		} else if strings.TrimSpace( string(request[:readLen]) ) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			// 反馈当前的时间信息
			conn.Write([]byte(daytime))
		}

		// 每次最后清理 request，否则因为 conn.Read() 会将新读取到的内容 append 到原内容之后
		request = make([]byte, 128)
	}


}

// 处理错误
func checkError(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
