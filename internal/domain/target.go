package domain

import (
	"math/rand"
	"net"
)

type Target struct {
	Ip   net.IP
	Port int
}

func (t *Target) GenerateRandomIP() {
	ip := make(net.IP, 4)

	ip[0] = byte(rand.Intn(256))
	ip[1] = byte(rand.Intn(256))
	ip[2] = byte(rand.Intn(256))
	ip[3] = byte(rand.Intn(256))

	t.Ip = ip
}
