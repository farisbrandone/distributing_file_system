package p2p

import "errors"

//ErrorInvalidHandshake is returned if the handleshake between
// the local and remote node could not be established.
 var ErrInvalidHanshake =errors.New("invalid hanshake")
//HanshakeFunc
type HanshakeFunc func(Peer)error

func NOPHanshakeFunc(Peer) error {return nil}