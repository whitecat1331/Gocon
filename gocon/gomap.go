package gocon

import (
	"context"
	"fmt"
	"log"

	"github.com/Ullaakut/nmap/v3"
)

func Gomap(addresses []string) {
	// Equivalent to
	// nmap -sV -T4 192.168.0.0/24 with a filter to remove non-RTSP ports.
	scanner, err := nmap.NewScanner(
		context.Background(),
		nmap.WithTargets(addresses),
		nmap.WithPorts("80", "554", "8554"),
		nmap.WithServiceInfo(),
		nmap.WithTimingTemplate(nmap.TimingAggressive),
		// Filter out ports that are not RTSP
		nmap.WithFilterPort(func(p nmap.Port) bool {
			return p.Service.Name == "rtsp"
		}),
		// Filter out hosts that don't have any open ports
		nmap.WithFilterHost(func(h nmap.Host) bool {
			// Filter out hosts with no open ports.
			for idx := range h.Ports {
				if h.Ports[idx].Status() == "open" {
					return true
				}
			}

			return false
		}),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		log.Printf("run finished with warnings: %s\n", *warnings) // Warnings are non-critical errors from nmap.
	}
	if err != nil {
		log.Fatalf("nmap scan failed: %v", err)
	}

	for _, host := range result.Hosts {
		fmt.Printf("Host %s\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d open with RTSP service\n", port.ID)
		}
	}
}
