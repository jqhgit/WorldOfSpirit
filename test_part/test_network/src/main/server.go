package main

import (
	"net"
	. "network"
	. "../logSys"
)
func HandleClientConnection(conn net.Conn, leave HandleDisconnect) {
	defer leave(conn)

	remote := conn.RemoteAddr()
	for {
		//test receive
		buf := make([]byte,1024)
		len,err := conn.Read(buf)
		if err != nil {
			LogDebug(err.Error())
			break
		}
		if len <= 0 {
			//client disconnect
			break
		}
		text := string(buf[:len])
		LogDebug("receive form ",remote," : ", text)

		//send
		text = "hello server received the message : " + text
		conn.Write([]byte(text))
	}
}

func HandleClientDisconnection(conn net.Conn) {
	defer conn.Close()

	LogDebug("disconnect remote:",conn.RemoteAddr())
}
func main() {
	var netImp NormalNetwork
	netImp.Listen(TCP, 44352, HandleClientConnection, HandleClientDisconnection)
}
