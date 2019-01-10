package main

import (
	"fmt"
	. "logSys"
	"net"
	. "network"
	"time"
)

func write_loop_string(conn net.Conn, ch chan string) {
	for {
		text := <- ch
		_,err := conn.Write([]byte(text))
		if err != nil {
			break
		}
		LogDebug("send to ",conn.RemoteAddr()," : ", text)
	}

	LogDebug("disconnected ...")
	conn.Close()
}

func read_loop_string(conn net.Conn) {
	for {
		buf := make([]byte,1024)
		len,err := conn.Read(buf)
		if err != nil {
			break
		}
		if buf == nil || len <= 0 {
			break
		}
		var str string = string(buf[:len])
		LogDebug("receive from ",conn.RemoteAddr()," : ", str)
	}
}


func main (){

	var netImp Net_normal
	conn,err := netImp.Connect(TCP, "127.0.0.1", 44352)
	if nil != err {
		LogDebug(err.Error())
		return
	}

	ch := make(chan string)
	go write_loop_string(conn, ch)
	go read_loop_string(conn)

	for {
		time.Sleep(1000)
		var sendText string
		fmt.Print("Enter what you need to send to sever:")
		fmt.Scanln(&sendText)
		if sendText == "exit" {
			break
		}
		size,err := conn.Write([]byte(sendText))
		if err != nil {
			LogDebug(err.Error())
			break
		}
		LogDebug("send size:", size)
	}
}
