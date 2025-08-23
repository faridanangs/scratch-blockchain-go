package main

import (
	"time"

	"github.com/faridanangs/projectx/network"
)

func main() {
	LocalTra := network.NewLocalTransport("LOCAL")
	RemoteTra := network.NewLocalTransport("REMOTE")

	RemoteTra.Connect(LocalTra)

	go func() {
		for {
			RemoteTra.SendMessage(LocalTra.Addr(), []byte("Hello Local Address"))
			time.Sleep(1 * time.Second)
		}
	}()

	serverOpts := network.ServerOpts{
		Transports: []network.Transport{
			LocalTra,
		},
	}

	s := network.NewServer(serverOpts)
	s.Start()

}
