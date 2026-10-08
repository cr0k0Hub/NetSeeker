package scanner

import (
	"NetSeeker/internal/domain"
	"net"
)

type Scanner struct {
	d      net.Dialer
	Result domain.TotalResult
}

func InitScanner(dial net.Dialer) *Scanner {
	return &Scanner{
		d:      dial,
		Result: domain.TotalResult{},
	}
}
