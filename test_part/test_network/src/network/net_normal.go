package network

import (
	. "logSys"
	"net"
	"strconv"
)

type Net_normal struct {

}

func (normal *Net_normal) Listen(network_type int, port int) error {
	var network = networkconv.ToString(network_type)
	var address = "localhost:"+strconv.Itoa(port)

	//begin to listen
	netListener,err := net.Listen(network, address)
	if err != nil {
		return err
	}
	LogDebug("begin to listen")

	//make sure listener to exit
	defer netListener.Close()

	//waiting for clients
	LogDebug("waiting for clients ...")
	for
	{
		conn,err := netListener.Accept()
		if err != nil {
			LogDebug(err.Error())
			continue
		}

		//client connected
		LogDebug(conn.RemoteAddr().String(), " ", network, " connect success")
		go normal.HnaldCilentConnection(conn)
	}

	return nil
}

func (normal *Net_normal) Connect(network_type int, ip string, port int) (net.Conn, error) {
	var network = networkconv.ToString(network_type)
	var address = ip+":"+strconv.Itoa(port)

	//begin to dial
	LogDebug("try to dial ", address, "...")
	conn,err := net.Dial(network, address)
	if err != nil {
		return nil,err
	}

	//success
	LogDebug("dial to ", address, " successfully.")

	return conn, nil
}

func (normal *Net_normal) HnaldCilentConnection(conn net.Conn) {
	defer conn.Close()

	remote := conn.RemoteAddr()
	for {
		//test receive
		buf := make([]byte,1024)
		len,err := conn.Read(buf)
		if err != nil {
			break
		}
		if buf == nil || len <= 0 {
			break
		}
		text := string(buf[:len])
		LogDebug("receive form ",remote," : ", text)

		//send
		text = "hello server received the message : " + text
		conn.Write([]byte(text))
	}

	LogDebug("disconnect remote:",remote)
}
