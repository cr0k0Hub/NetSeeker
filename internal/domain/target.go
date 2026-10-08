package domain

import "net"

type Target struct {
	Ip   net.IP
	Port int
}
