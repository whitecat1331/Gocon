package gocon

import (
	"fmt"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/likexian/whois"
	"github.com/prometheus-community/pro-bing"

	"github.com/whitecat1331/godevsuite"
	gowitness "github.com/whitecat1331/gowitness/cmd"
	"github.com/whitecat1331/httprobe"
)

// domain scraper

type DomainScraper struct {
	InitialDomains []string
	AssetFinder    []string
	AmassEnum      []string
	WaybackURLS    []string
	ScrapedDomains []string
}

var domainScraperLog = setLogName("domain_scraper.log")

func NewDomainScraper(initialDomains []string) *DomainScraper {

	logger, f := godevsuite.SetupSLogger(domainScraperLog, false, nil)
	defer f()
	domainScraper := new(DomainScraper)
	domainScraper.InitialDomains = initialDomains
	assetFinderResults, err := FetchAssetFinder(domainScraper.InitialDomains)
	if err != nil {
		logger.Error("")
	}
	domainScraper.AssetFinder = assetFinderResults
	domainScraper.AmassEnum = FetchAmassEnum(initialDomains)

	var scrapedDomains []string
	scrapedDomains = append(scrapedDomains, domainScraper.AssetFinder...)
	scrapedDomains = append(scrapedDomains, domainScraper.AmassEnum...)

	domainScraper.ScrapedDomains = scrapedDomains

	return domainScraper
}

// domain enumeration
func FetchWhois(domain string) string {
	result, err := whois.Whois(domain)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// probes
func FetchHTTProbe(domains []string) []string {
	return httprobe.HTTProbe(domains)
}

func ActivePing(domain string, count int, timeout int) bool {
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

func FetchActive(domains []string) []string {
	return httprobe.HTTProbe(domains)
}

func fetchGoWitness(urls []string, threads int) {
	gowitness.GoWitnessess(urls, threads)
}

func FetchGoWitnessDefault(urls []string) {
	fetchGoWitness(urls, 4)
}

type GoCon struct {
	AssetFinder []string
}

func NewGoCon(domains []string) *GoCon {
	return &GoCon{
		AssetFinder: FetchAssetFinder(domains),
	}
}
