package main

import (
	"fmt"
	"os"

	"github.com/hatt3r/attack-surface-mapper/internal/fingerprint"
	"github.com/hatt3r/attack-surface-mapper/internal/risk"
	"github.com/hatt3r/attack-surface-mapper/internal/scanner"
)

func main() {

	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Usage: scanner <target-ip>")
		return
	}

	target := os.Args[1]

	fmt.Println("Scanning target:")
	fmt.Println("Target:", target)

	openPorts := scanner.ScanTarget(target)

	fmt.Println("Open Portsdetected \n:")
	for _, port := range openPorts {
		service := fingerprint.IdentifyService(port)
		risklevel := risk.Score(service)
		fmt.Printf("Port: %d -> %s, Risk Level: %s\n", port, service, risklevel)
	}
}
