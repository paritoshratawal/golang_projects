package p2p

import "net"

// Peer is an interface that represents remote node
type Peer interface {
	// close() error
}

/*
Transport handles communication between two nodes in the network,
this can be from (TCP, UDP, websockets....)
*/
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}

// TCPPeer represents the TCP node over a TCP established connection
type TCPPeer struct {
	conn net.Conn

	// if we dial and retrieve a conn outbound = true
	// if we accept and retrieve a conn outbound = false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

func OnPeer(peer Peer) error {
	return nil
}
