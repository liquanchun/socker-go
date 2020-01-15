package main

import (
	"net"
	"os"
	"socker/gb"
	logCondf "socker/gb/conf"
)

func init() {
	logCondf.LogConf()
}

func main() {
	//建立socket端口监听
	netListen, err := net.Listen("tcp", "193.112.155.251:8080")
	CheckError(err)

	defer netListen.Close()

	gb.Logger.Info("Waiting for clients to connetion port 8080......")

	//等待客户端访问
	for {
		conn, err := netListen.Accept() //监听接收
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

//错误处理
func CheckError(err error) {
	if err != nil {
		gb.Logger.Error(os.Stderr, "Fatal error: %s", err.Error())
	}
}
