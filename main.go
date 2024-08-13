package main

import (
	"fmt"
	"log"

	"github.com/farisbrandone/distributed_file_storage/p2p"
)

func OnPeer(peer p2p.Peer)error{
	//fmt.Println("doing same log with the peer outside of TCPTransport")
	/*return fmt.Errorf("failed the onpeer function")*/
	peer.Close()
	return nil
}

func main(){
	tcpOpts:=p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		HanshakeFunc: p2p.NOPHanshakeFunc ,
		Decoder: p2p.DefaultDecoder{},
		OnPeer: OnPeer,
	}
	tr :=p2p.NewTCPTransport(tcpOpts)
    
  go func(){
	for {
		msg:= <-tr.Consume()
		fmt.Printf("%+v\n", msg)
	}
	
  }()

   if err:=tr.ListenAndAccept(); err!=nil {
	log.Fatal(tr.ListenAndAccept())
   }
   
  /*ci:=p2p.Couscous{
	Al: 10,
	Bl: 11,
  }
  d:=ci.Blabla(12,10)
	fmt.Println("We good")
	fmt.Printf("VALUE OF CI %+v\n ",ci)
	fmt.Printf("VALUE OF d %+v\n ",d)*/
	select {}
}