package server

import "net"

func Server() {
	net.Listen("tcp", "localhost:8080")
}
