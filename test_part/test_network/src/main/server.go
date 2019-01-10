package main

import (
	. "network"
)

func main() {
	var netImp Net_normal
	netImp.Listen(TCP, 44352)
}
