package output

import (
	"NetSeeker/internal/domain"
	"fmt"
	"time"
)

func PrintLogo() {
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

func PrintInfo(n int, p int, t int, timeout time.Duration) {
	fmt.Printf("NetSeeker v0.1\n")
	fmt.Printf("Port scanner\n")
	fmt.Printf("\nTarget: random IPv4\n")
	fmt.Printf("Targets: %d\n", n)
	fmt.Printf("Port:    %d\n", p)
	fmt.Printf("Workers: %d\n", t)
	fmt.Printf("Timeout: %v\n", timeout)
	fmt.Printf("\n--------------------------------------------------------------------------------\n\n")
}

func PrintStats(t *domain.TotalResult) {
	fmt.Printf("Scanned:     %d\n", t.Scanned.Load())
	fmt.Printf("Opened:      %d\n", t.Opened.Load())
	fmt.Printf("Closed:      %d\n", t.Closed.Load())
	fmt.Printf("Timeout:     %d\n", t.Timeout.Load())
	fmt.Printf("Unreachable: %d\n", t.Unreachable.Load())
	fmt.Printf("Errors:      %d\n", t.Errors.Load())
}
