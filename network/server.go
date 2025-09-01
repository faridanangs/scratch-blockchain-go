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
	s.initialTransports()
	t := time.NewTicker(time.Second * 5)

Free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break Free
		case <-t.C:
			fmt.Println("Do stuff every 5 seconds")
		}
	}
	fmt.Println("Shutdown")

}

func (s *Server) initialTransports() {
	for _, tra := range s.Transports {
		go func(tra Transport) {
			for rpc := range tra.Consume() {
				s.rpcCh <- rpc
			}
		}(tra)
	}
}
