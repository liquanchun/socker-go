package main

import (
	"fmt"
	"net"
	"os"
	"socker-go/gb"
	logCondf "socker-go/gb/conf"
)

//与服务器相关的资源都放在这里面
type TcpServer struct {
	listener   *net.TCPListener
	hawkServer *net.TCPAddr
}

func init() {
	logCondf.LogConf()
}

func main() {
	hawkServer, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		gb.Logger.Error(os.Stderr, "Fatal error: %s", err.Error())
		fmt.Printf("Fatal error: %s", err.Error())
		return
	}
	//建立socket端口监听
	netListen, err := net.ListenTCP("tcp", hawkServer)
	if err != nil {
		gb.Logger.Error(os.Stderr, "Fatal error: %s", err.Error())
		fmt.Printf("Fatal error: %s", err.Error())
		return
	}

	defer netListen.Close()
	tcpServer := &TcpServer{
		listener:   netListen,
		hawkServer: hawkServer,
	}
	gb.Logger.Info("Waiting for clients to connetion port 8080......")

	//等待客户端访问
	for {
		conn, err := tcpServer.listener.Accept() //监听接收
		if err != nil {
			continue //如果发生错误，继续下一个循环。
		}

		gb.Logger.Info(conn.RemoteAddr().String(), "tcp connect success") //tcp连接成功
		go handleConnection(conn)
	}

}

//处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048) //建立一个slice
	for {
		n, err := conn.Read(buffer) //读取客户端传来的内容
		if err != nil {
			gb.Logger.Error(conn.RemoteAddr().String(), "connection error: ", err)
			return //当远程客户端连接发生错误（断开）后，终止此协程。
		}

		gb.Logger.Info(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

		//返回给客户端的信息
		strTemp := "send data is ok : " + string(buffer[:n])
		conn.Write([]byte(strTemp))
	}
}
