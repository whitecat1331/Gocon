package gocon

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/haccer/subjack/subjack"
)

func Subjack(subdomain string) bool {
	var fingerprints []subjack.Fingerprints
	config, _ := os.ReadFile("custom_fingerprints.json")
	err := json.Unmarshal(config, &fingerprints)
	if err != nil {
		log.Fatal(err)
	}

	/* Use subjack's advanced detection to identify
	if the subdomain is able to be taken over. */
	service := subjack.Identify(subdomain, false, false, 10, fingerprints)

	if service != "" {
		service = strings.ToLower(service)
		log.Printf("%s is pointing to a vulnerable %s service.\n", subdomain, service)
	}

	return service != ""
}
