package network

import "net"

//network type define
const (
	TCP = iota
	TCP4
	TCP6
	UDP
	UDP4
	UDP6
	IP
	IP4
	IP6
)

// Callback function while remote disconnect
type HandleDisconnect func(conn net.Conn)

// Callback function while remote connect
type HandleConnect func(conn net.Conn, leave HandleDisconnect)

// The interface of network implement
type NetInterface interface {

	// Listen announces on the local network address, you can use it for the server
	// to receive remote request.
	//
	// The network type must be TCP TCP4 TCP6, and you need to set the connect and
	// disconnect callback function before you call function "Listen".
	//
	// When remote connect, the function "HandleConnect" will be called as goroutine,
	// and when remote disconnect, the function "HandleDisconnect" will be called in
	// function "HandleConnect".
	Listen(networkType int, port int, connect HandleConnect, disconnect HandleDisconnect) error

	// Connect to the address on the named network, if you want to connect to a known
	// remote server you can use it.
	//
	// The network type must be TCP TCP4 TCP6 UDP UDP4 UDP6, and the function wile return
	// the connection implement and the error when connect to remote.
	Connect(networkType int, ip string, port int) (net.Conn, error)
}


// converter for int type network to string
var networkconv ConvertNetwork

type ConvertNetwork struct {
}

func (t *ConvertNetwork) ToString(network int) string {
	switch
	{
	case network == TCP:
		return "tcp"
	case network == UDP:
		return "udp"
	default:
		return "error"
	}
}


