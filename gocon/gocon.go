package gocon

import (
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/prometheus-community/pro-bing"

	"github.com/whitecat1331/amass/v4/amass"
	"github.com/whitecat1331/assetfinder"
	"github.com/whitecat1331/httprobe"
	"github.com/whitecat1331/waybackurls"
	// "github.com/whitecat1331/gowitness"
)

func fetchAssetFinder(domain string) []string {
	return assetfinder.AssetFinder(domain, true)
}

func fetchWaybackURLs(domains []string) []string {
	return waybackurls.WaybackURLS(domains)
}

// might remove output option
func fetchAmassEnum(domain string, output string) []string {
	return amass.EnumAllDomain(domain, output)
}

func fetchHTTProbe(domains []string) []string {
	return httprobe.HTTProbe(domains)
}

func activePing(domain string, count int, timeout int) bool {
	pinger, err := probing.NewPinger(domain)
	if err != nil {
		log.Fatal(err)
	}

	pinger.Count = count
	pinger.Timeout = time.Duration(timeout) * time.Second
	err = pinger.Run()
	if err != nil {
		log.Fatal(err)
	}

	stats := pinger.Statistics()

	return stats.PacketsRecv > 0

}

func fetchSubjack(domains []string) []string {
	var potentialTakeovers []string
	for _, domain := range domains {
		if Subjack(domain) {
			potentialTakeovers = append(potentialTakeovers, domain)
		}
	}

	return potentialTakeovers
}

func fetchActive(domains []string) []string {
	return httprobe.HTTProbe(domains)
}

func fetchGoWitness(urls string) {
	// gowitness.GoWitness()
}
