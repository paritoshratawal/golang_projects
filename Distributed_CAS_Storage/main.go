package main

import (
	"fmt"
	"log"

	"github.com/paritoshratawal/distributed_cas_storage/p2p"
)

func main() {

	tcpOpts := p2p.TCPOptions{
		ListenAddr: ":4040",
		Handshake:  p2p.NOHandshake,
		Decoder:    p2p.DefaultDecoder{},
		OnPeer:     p2p.OnPeer,
	}
	transport := p2p.NewTCPTransport(tcpOpts)
	if err := transport.ListenAndAccept(); err != nil {
		log.Fatal("Error while Transport ListenAndAccept()", err)
	}

	go func() {
		for {
			msg := <-transport.Consume()
			fmt.Printf("Message: %v\n", msg)
		}
	}()

	select {}
}
