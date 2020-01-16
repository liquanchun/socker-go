package main

import (
	"fmt"
	"net"
	"os"
	"socker-go/gb"
	logCondf "socker-go/gb/conf"
	"strconv"
)

const (
	SERVER_IP       = "0.0.0.0"
	SERVER_PORT     = 8080
	SERVER_RECV_LEN = 1024
)

func init() {
	logCondf.LogConf()
}
func main() {
	address := SERVER_IP + ":" + strconv.Itoa(SERVER_PORT)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer listener.Close()
	gb.Logger.Info("Waiting for clients to connetion port 8080......")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer conn.Close()
		for {
			data := make([]byte, SERVER_RECV_LEN)
			_, err := conn.Read(data)
			if err != nil {
				fmt.Println(err)
				break
			}

			var slice []byte
			for _, d := range data {
				if d > 0 {
					slice = append(slice, d)
				}
			}
			strData := string(slice)
			gb.Logger.Info(conn.RemoteAddr().String(), "receive [", strData, "]")
			_, err = conn.Write([]byte("ok"))
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}
