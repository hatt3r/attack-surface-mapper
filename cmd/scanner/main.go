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
	//calling scanner and scantarget function to get ports

	fmt.Println("Open Ports detected:\n")
	for _, port := range openPorts {
		service := fingerprint.IdentifyService(port)
		//calling fingerprint and identifyservice function to get service name used

		risklevel := risk.Score(service)
		//calling risk and score function to get predefined risk level for the service

		fmt.Printf("Port: %d -> %s, Risk Level: %s\n", port, service, risklevel)
	}
}
