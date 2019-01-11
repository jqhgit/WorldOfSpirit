package network

import (
	. "logSys"
	"net"
	"strconv"
)

type NormalNetwork struct {

}

func (normal *NormalNetwork) Listen(networkType int, port int, connect HandleConnect, disconnect HandleDisconnect) error {
	var network = networkconv.ToString(networkType)
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
		go connect(conn, disconnect)
	}

	return nil
}

func (normal *NormalNetwork) Connect(networkType int, ip string, port int) (net.Conn, error) {
	var network = networkconv.ToString(networkType)
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
