package scanner

import (
	"NetSeeker/internal/domain"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"syscall"
)

// TODO: rewrite scanner. Targets creates workers and workers do scan shit
func generateRandomIP() net.IP {
	ip := make(net.IP, 4)

	ip[0] = byte(rand.Intn(256))
	ip[1] = byte(rand.Intn(256))
	ip[2] = byte(rand.Intn(256))
	ip[3] = byte(rand.Intn(256))

	return ip
}

func (s *Scanner) defineStatus(err error) string {
	if errors.Is(err, syscall.ECONNREFUSED) {
		s.Result.Closed.Add(1)
		return "closed"
	}

	if errors.Is(err, syscall.ECONNRESET) {
		s.Result.Closed.Add(1)
		return "closed"
	}

	if errors.Is(err, syscall.ENETUNREACH) {
		s.Result.Unreachable.Add(1)
		return "unreachable"
	}

	if errors.Is(err, syscall.EHOSTUNREACH) {
		s.Result.Unreachable.Add(1)
		return "unreachable"

	}

	if errors.Is(err, syscall.ENETDOWN) {
		s.Result.Unreachable.Add(1)
		return "unreachable"
	}

	s.Result.Errors.Add(1)
	return "error"
}

func (s *Scanner) Scan(ctx context.Context, target domain.Target) domain.ScanResult {
	addr := target.Ip.String() + ":" + strconv.Itoa(target.Port)
	conn, err := s.d.DialContext(ctx, "tcp", addr)
	//s.Result.Scanned.Add(1)
	// TODO: fix stats count
	s.Result.Scanned.Add(1)
	if err != nil {
		return domain.ScanResult{
			Target: target,
			Status: s.defineStatus(err),
		}
	}

	defer conn.Close()

	s.Result.Opened.Add(1)
	return domain.ScanResult{
		Target: target,
		Status: "opened",
	}
}

func (s *Scanner) Worker(ctx context.Context, portsOpen chan<- string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ip := generateRandomIP()
			target := domain.Target{
				Ip:   ip,
				Port: port,
			}
			result := s.Scan(ctx, target)
			//fmt.Printf("Scanning IP %s:%d STATUS: %s\n", target.Ip, target.Port, result.Status)
			if result.Status == "opened" {
				select {
				case portsOpen <- fmt.Sprintf("[+] %s:%d OPENED!", target.Ip, target.Port):
				case <-ctx.Done():
					return
				}
			}
		}
	}
}
