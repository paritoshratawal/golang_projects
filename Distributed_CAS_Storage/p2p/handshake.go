package p2p

// import "errors"

// var ErrInvalidHanshake = errors.New("Invalid handshake")

type HandshakeFunc func(Peer) error

func NOHandshake(Peer) error { return nil }
