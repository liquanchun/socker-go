package main

import (
	"fmt"
	"github.com/tidwall/evio"
	"os"
	"socker-go/gb"
	logCondf "socker-go/gb/conf"
)

func init() {
	logCondf.LogConf()
	gb.InitDb()
}

// 添加git分支
func main() {
	var events evio.Events
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		out = in
		gb.Logger.Info("IP:", c.RemoteAddr().String(), "，receive[", string(in), "]")
		gb.SaveToDB(string(in), c.LocalAddr().String())
		fmt.Println(string(in))
		return
	}
	if err := evio.Serve(events, "tcp://0.0.0.0:8080"); err != nil {
		gb.Logger.Error(os.Stderr, "Fatal error: %s", err.Error())
		panic(err.Error())
	} else {
		gb.Logger.Info("Waiting for clients to connetion port 8080......")
	}
}
func Test() {
	fmt.Println("this is test func")
	fmt.Println("add test branch")
	fmt.Println("add abc branch")
}
