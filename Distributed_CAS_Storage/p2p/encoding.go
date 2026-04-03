package p2p

import (
	"encoding/gob"
	"fmt"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader, msg *RPC) error {
	return gob.NewDecoder(r).Decode(msg) //Reading from network
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, msg *RPC) error {
	fmt.Println("Decoding Incoming message")
	buffer := make([]byte, 1028)
	n, err := r.Read(buffer)
	if err != nil {
		fmt.Printf("TCP error %v\n", err)
	}
	msg.Payload = buffer[:n]

	return nil
}
