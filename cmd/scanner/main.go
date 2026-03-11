package main

import (
	"fmt"
	"os"

	"github.com/hatt3r/attack-surface-mapper/internal/scanner"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: scanner <target-ip>")
		return
	}

	target := os.Args[1]

	fmt.Println("Scanning target:", target)

	openPorts := scanner.ScanTarget(target)

	fmt.Println("Open Ports:")
	for _, port := range openPorts {
		fmt.Println(port)
	}
}
