package p2p

import (
	//"bytes"
	"encoding/gob"
	//"fmt"
	"io"
)

type Decoder interface{
	Decode(io.Reader, *RPC)error
}

type GOBDecoder struct{}

type DefaultDecoder struct{}

type Couscous struct{
	Al float32
	Bl float32
}

func (dec GOBDecoder) Decode (r io.Reader, msg *RPC)error {
	return gob.NewDecoder(r).Decode(msg)
}

/*func (dec DefaultDecoder) Decode (r io.Reader, msg *Message)error {
	return gob.NewDecoder(r).Decode(msg)
}*/

type Coco interface {
	Blabla(string)
}

func (ci *Couscous)Blabla(a float32, b float32)float64{
  ci.Al+=a
  ci.Bl+=b
return float64(ci.Al)+float64(ci.Bl)
  
}

type NOPDecode struct{}

func (dec DefaultDecoder)Decode(r io.Reader, msg *RPC)error {
	buf:=make([]byte, 1028)
	n, err := r.Read(buf)
	if err !=nil {
		return err
	}
	msg.Payload=buf[:n]

	return nil
}