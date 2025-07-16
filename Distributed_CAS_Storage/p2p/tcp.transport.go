package p2p

import (
	"fmt"
	"net"
)

type TCPOptions struct {
	ListenAddr string
	Handshake  HandshakeFunc
	Decoder    Decoder //Interface is implemented here
	OnPeer     func(Peer) error
}

type TCPTransport struct {
	TCPOptions
	Listener net.Listener
	RPCChan  chan RPC

	// mutex sync.RWMutex
	// peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPOptions) *TCPTransport {
	return &TCPTransport{
		TCPOptions: opts,
		RPCChan:    make(chan RPC),
	}
}

// For reading incoming messages recieving from another peer
func (t *TCPTransport) Consume() <-chan RPC {

	return t.RPCChan
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddr)

	if err != nil {
		return err
	}

	fmt.Println("TCP server listening to port", t.ListenAddr)
	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	fmt.Println("started Accept Loop")
	for {
		fmt.Println("started FOR loop")
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		fmt.Printf("New peer connection %v\n", conn)
		go t.handleConn(conn)
	}
}

// type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Println("Handle Connection")

	var err error
	defer func() {
		fmt.Println("Droping peer connection", err)
		conn.Close()
	}()
	peer := NewTCPPeer(conn, true)

	if err := t.Handshake(peer); err != nil {
		fmt.Printf("Handshaking error %v\n", err)
		return
	}
	fmt.Println("Handshake Success")

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}

	// buffer := make([]byte, 2000)
	rpc := RPC{}
	for {
		fmt.Println("Start Recieving Data From Client")
		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP error %v\n", err)
			continue
		}
		rpc.From = conn.RemoteAddr()
		t.RPCChan <- rpc
		fmt.Printf("Message: %v\n", rpc)
	}
}
