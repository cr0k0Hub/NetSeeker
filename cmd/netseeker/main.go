package main

import (
	"NetSeeker/internal/scanner"
	"context"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func printLogo() {
	fmt.Printf(
		"  ▄        ▄  ▄                 ▄      ▄  ▄     ▄  ▄     ▄    ▄   ▄  ▄     ▄  ▄  \n" +
			"▄██ ▄█▄  ▄██▄▀██  ▀███  ███▀  ▄▀██▄  ▄██▄▀██  ▄██▄▀██  ▄██  ▄██ ▄██▄▀██  ▄██▄▀██ \n" +
			" █░▀ ▀░█  █░   ▀    ▀  ░ ▀   █░  ▀██  █░   ▀   █░   ▀   █░   ░█  █░   ▀   █░   ░█\n" +
			" ▒▓   ▒▓  ▒▓          ▒▓     ▒▓   ▀▀  ▒▓       ▒▓       ▒▓   ▒   ▒▓       ▒▓   ▒ \n" +
			" ▀▀   ▀▀  ▀▀▀▀▀       ▀▀      ▀▀      ▀▀▀▀▀    ▀▀▀▀▀    ▀▀▀▀▀    ▀▀▀▀▀    ▀▀▀▀▀  \n" +
			" ░░   ░░  ░░          ░░        ▀░▄   ░░       ░░       ░░   ░   ░░       ░░  ░  \n" +
			" ▒▒   ▒▒  ▒▒          ▒▒          ▒▒  ▒▒       ▒▒       ▒▒   ▒▒  ▒▒       ▒▒  ▒▒ \n" +
			" ▓▓   ▓▓  ▓▓▄  ▓▓     ▓▓▄    ▓▓▄  ▓▓  ▓▓▄  ▓▓  ▓▓▄  ▓▓  ▓▓   ▓▓  ▓▓▄  ▓▓  ▓▓  ▓▓ \n" +
			"▀██  ▀██   ▀██▄▀      █▀      ▀██▄▀    ▀██▄▀    ▀██▄▀  ▀██  ▀██   ▀██▄▀  ▀██ ▀██ \n" +
			"  ▀    ▀     ▀       ▀          ▀        ▀        ▀      ▀    ▀     ▀      ▀   ▀ \n")
	fmt.Println("                                                                        by prikin")
}

// TODO: create Output domain and place all print shit into it
func printInfo(n int, p int, t int, timeout time.Duration) {
	fmt.Printf("NetSeeker v0.1\n")
	fmt.Printf("Port scanner\n")
	fmt.Printf("\nTarget: random IPv4\n")
	fmt.Printf("Targets: %d\n", n)
	fmt.Printf("Port:    %d\n", p)
	fmt.Printf("Workers: %d\n", t)
	fmt.Printf("Timeout: %v\n", timeout)
	fmt.Printf("\n--------------------------------------------------------------------------------\n\n")
}

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

	printLogo()
	printInfo(*numIPs, *port, *numThreads, *timeout)

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
	s.Result.Print()
}
