package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts
	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)

Free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break Free
		case <-ticker.C:
			fmt.Println("do stuf every x second")
		}

	}
	fmt.Println("Server Shutdown")

}

func (s *Server) initTransports() {
	for _, tra := range s.Transports {
		go func(tra Transport) {
			for rpc := range tra.Consume() {
				s.rpcCh <- rpc
			}
		}(tra)
	}
}
