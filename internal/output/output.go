package output

import (
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
