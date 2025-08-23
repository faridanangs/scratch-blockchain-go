package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.RWMutex
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (lt *LocalTransport) Consume() <-chan RPC {
	return lt.consumeCh
}

func (lt *LocalTransport) Connect(tra Transport) error {
	lt.lock.Lock()
	defer lt.lock.Unlock()

	lt.peers[tra.Addr()] = tra.(*LocalTransport)

	return nil
}

func (lt *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	lt.lock.RLock()
	defer lt.lock.RUnlock()

	peer, ok := lt.peers[to]
	if !ok {
		return fmt.Errorf("%s: Could not send message to %s", lt.addr, to)
	}

	peer.consumeCh <- RPC{
		From:    lt.addr,
		Payload: payload,
	}

	return nil
}

func (lt *LocalTransport) Addr() NetAddr {
	return lt.addr
}
