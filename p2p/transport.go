package p2p

// Peer is the interface that representats the remote node
type Peer interface {
    Close()error
}

// transport is anything that handles the communication
// between the nodes in the network
// form (tcp, UDP, websockets, ....)
type Transport interface{
	ListenAndAccept() error
	Consume() <-chan RPC
}