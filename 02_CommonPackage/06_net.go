package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 创建一个tcp的监听器，监听本地的端口8080
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on 127.0.0.1:8080")
	for {
		// 等待连接，当有新的连接时，Accept函数会返回一个Conn对象，表示客户端与服务器之间的连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Accepted new connection from:", conn.RemoteAddr().String())

		// 开启一个新的协程来处理客户端请求
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// 当函数结束时关闭连接
	defer conn.Close()

	// 读取客户端发送的请求
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	// 将请求的数据转换成字符串，并打印输出
	request := string(buf[:n])
	fmt.Println("Received request from client:", request)

	// 构造响应数据并返回给客户端
	response := "Hello, client!"
	conn.Write([]byte(response))
	fmt.Println("Sent response to client:", response)
}
