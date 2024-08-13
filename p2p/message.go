package p2p

import "net"

//Message holds any arbitrary data that is being sent over the
//each transport between two nodes to the networks
type RPC struct {
	From net.Addr
	Payload []byte
}