package network

type NetAddress string

type RPC struct {
	From    NetAddress
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(tra Transport) error
	SendMessage(to NetAddress, msg []byte) error
	Address() NetAddress
}
