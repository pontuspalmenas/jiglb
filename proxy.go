package main

import (
	"net"
	"sync"
	"time"
)

type Proxy struct {
	ListenAddress string
	Connections   map[string]*ClientConnection
	ConnMutex     sync.Mutex
}

type ClientConnection struct {
	ClientAddr     net.Addr
	Conn           net.Conn
	TargetConn     net.Conn
	LastActiveTime time.Time
	Buffer         chan []byte
}
