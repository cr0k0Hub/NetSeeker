package main

import (
	"NetSeeker/internal/output"
	"NetSeeker/internal/scanner"
	"context"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	numIPs := flag.Int("n", 1, "Number of IP addresses to gather")
	port := flag.Int("p", 80, "Port to check if open")
	numThreads := flag.Int("t", 10, "Number of threads to use")
	timeout := flag.Duration("timeout", 3*time.Second, "Timeout")

	flag.Parse()

	portsOpen := make(chan string, *numIPs)
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	d := net.Dialer{Timeout: *timeout}
	s := scanner.InitScanner(d)

	output.PrintLogo()
	output.PrintInfo(*numIPs, *port, *numThreads, *timeout)

	for i := 0; i < *numThreads; i++ {
		wg.Add(1)
		go s.Worker(ctx, portsOpen, *port, &wg)
	}

	for i := 0; i < *numIPs; i++ {
		fmt.Println(<-portsOpen)
	}

	cancel()
	wg.Wait()
	close(portsOpen)
	fmt.Printf("\n--------------------------------------------------------------------------------\n\n")
	output.PrintStats(&s.Result)
}
