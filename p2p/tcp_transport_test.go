package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T){
      
	cpOpts:=TCPTransportOpts{
		ListenAddr: ":3000",
		HanshakeFunc: NOPHanshakeFunc ,
		Decoder: DefaultDecoder{},
	}


	listenAddr:= ":3000"
	tr:=NewTCPTransport(cpOpts)
	assert.Equal(t, tr.ListenAddr, listenAddr)

	//server
   assert.Nil(t,tr.ListenAndAccept()) 
}