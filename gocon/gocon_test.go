package gocon

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGocon(t *testing.T) {
	domain := os.Getenv("DOMAIN")
	output := os.Getenv("OUTPUT")
	domains := []string{domain}
	fmt.Printf("domain: %s\noutput: %s\ndomains: %s\n", domain, output, domains)

	// fmt.Println("AssetFinder")
	// assets := fetchAssetFinder(domain)
	// fmt.Println(assets)
	// fmt.Println("AmassEnum")
	// amassEnum := fetchAmassEnum(domain, output)
	// fmt.Println(amassEnum)
	// fmt.Println("SSLMate")
	// sslMate := fetchSSLMate(domain)
	// fmt.Println(sslMate)
	// fmt.Println("Ping")
	// res := activePing(domain, 3, 2)
	// fmt.Printf("%t\n", res)
	// fmt.Println("WebDomain")
	// webDomains := fetchActive(domains)
	// fmt.Println(webDomains)
	// fmt.Println("Takeovers")
	potentialTakeovers := fetchSubjack(domains)
	fmt.Println(potentialTakeovers)
	// fmt.Println("HTTProbe")
	// httpURLS := fetchHTTProbe(domains)
	// fmt.Println(httpURLS)
	// fmt.Println("WaybackURLs")
	// waybackurls := fetchWaybackURLs(domains)
	// fmt.Println(waybackurls)
	// fmt.Println("DNSInfo")
	// dnsInfo := fetchDNSInfo(domain, "8.8.8.8:53")
	// fmt.Println(dnsInfo)

}
