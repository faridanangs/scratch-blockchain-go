package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddress
	consumeCh chan RPC
	peers     map[NetAddress]*LocalTransport
	lock      sync.RWMutex
}

func NewLocalTransport(addr NetAddress) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddress]*LocalTransport),
	}
}

func (ltra *LocalTransport) Consume() <-chan RPC {
	return ltra.consumeCh
}

func (ltra *LocalTransport) Connect(tra Transport) error {
	ltra.lock.Lock()
	defer ltra.lock.Unlock()

	ltra.peers[tra.Address()] = tra.(*LocalTransport)

	return nil
}

func (ltra *LocalTransport) SendMessage(to NetAddress, msg []byte) error {
	ltra.lock.RLock()
	defer ltra.lock.RUnlock()

	peer, ok := ltra.peers[to]
	if !ok {
		return fmt.Errorf("%s: Could not send message to %s", ltra.addr, to)
	}

	peer.consumeCh <- RPC{
		From:    ltra.addr,
		Payload: msg,
	}

	return nil
}

func (ltra *LocalTransport) Address() NetAddress {
	return ltra.addr
}
