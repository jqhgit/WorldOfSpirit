package network

//network_type
const (
	TCP = iota
	UDP
)
type net_interface interface {

}

type network_conv struct {

}
func (t *network_conv) ToString(network int) string {
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
var networkconv network_conv

