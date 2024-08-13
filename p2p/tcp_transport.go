package p2p

import (
	//"bytes"
	"fmt"
	"net"
	"reflect"
	"sync"
)

//TCPPeer represents the remote node over a TCP established connection

type TCPPeer struct{
	//conn is the underlying connection of the peer
	conn net.Conn

	//if we dial a connection => outbound==true
	//if we accept and retrieve a conn => outbound==false
	outbound bool
}

type TCPTransportOpts struct {
	ListenAddr string
	HanshakeFunc HanshakeFunc
	Decoder Decoder
	OnPeer func(Peer)error
}

type TCPTransport struct {
	TCPTransportOpts
	//ListenAdress string
	Listener net.Listener
	//shakeHands HanshakeFunc
	//decoder Decoder
    rpcch chan RPC

	mu sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}

func (p *TCPPeer)Close()error {
	return p.conn.Close()
}
 
//consume implements the Transport interface, 
//which will return only channel for reading incoming received
//from another peer in the network
func (t *TCPTransport) Consume()<-chan RPC {
	return t.rpcch
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport {
		TCPTransportOpts: opts,
		rpcch: make(chan RPC),
		//shakeHands: NOPHanshakeFunc,
		//ListenAdress: listenAddr,
	}
}

func (t *TCPTransport)ListenAndAccept() error{
	var err error
	t.Listener,err =net.Listen("tcp", t.ListenAddr)
	//ln, err :=net.listen("tcp", t.ListenAdress)
	if err !=nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop(){
	for {
		conn,err := t.Listener.Accept()
		if err!=nil {
			fmt.Printf("TCP accept error:%s\n ", err)
		}

		go t.handleConn(conn)
	}
}

type Temp struct {}

func (t *TCPTransport)handleConn(conn net.Conn){
	 var err error
	

	defer func(){
       fmt.Printf("dropping connection : %s", err)
	   conn.Close()
	}()

	peer := NewTCPPeer(conn, true)

    if err:=t.HanshakeFunc(peer); err!=nil {
		/*conn.Close()
		fmt.Printf("TCP hanshake error : %+s\n", err)*/
		return
		//return ErrInvalidHanshake
	}

	if t.OnPeer!=nil {
		if err:=t.OnPeer(peer); err!=nil{
			return
		}
	}
 
	//buf:=new(bytes.Buffer)

	//Read Loop
	rpc:=RPC{}
	//buf:=make([]byte, 2000)
    ci:=Couscous{
		Al: 10,
		Bl: 11,
	  }
	  fmt.Printf("dondonca")
	for {
           // n,err :=conn.Read(buf)
      err:=t.Decoder.Decode(conn, &rpc)// on a d'abord 
	  fmt.Println(reflect.TypeOf(err))
	 // panic(err)
	  //fait la conndition sur error et ca affichais de manière continu
	 /* if err==net.ErrClosed{//on a utiliser ca pour empeecher le problème
		return
         
	 }*/
	 if err !=nil {
		fmt.Printf("TCP error : %+s\n", err)
		return
	 }
	 rpc.From=conn.RemoteAddr()
	 t.rpcch <-rpc
	 fmt.Printf("message : %+v\n", rpc/*buf[:n]*/)
	 
	  d:=ci.Blabla(12,10)
		fmt.Println("We good")
		fmt.Printf("VALUE OF CI %+v\n ",ci)
		fmt.Printf("VALUE OF d %+v\n ",d)
	}
        
   
	//fmt.Printf("new incoming connection %+v\n", peer)
}