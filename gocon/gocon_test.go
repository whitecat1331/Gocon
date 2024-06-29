package gocon

import (
	"os"
	"strings"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

type testValues struct {
	Domain  string
	Domains []string
	Output  string
}

func NewTestingValues() testValues {
	domains := strings.Split(os.Getenv("DOMAINS"), ",")
	return testValues{
		Domain:  domains[0],
		Domains: domains,
		Output:  os.Getenv("OUTPUT"),
	}

}

var tv = NewTestingValues()

func TestFetchAssetFinder(t *testing.T) {
	domains := FetchAssetFinder(tv.Domains)
	t.Log(domains)
}

func TestFetchAmassEnum(t *testing.T) {
	domains := FetchAmassEnum(tv.Domains)
	t.Log(domains)
}

func TestFetchWaybackURLs(t *testing.T) {
	domains := FetchWaybackURLs(tv.Domains)
	t.Log(domains)
}

// func TestGocon(t *testing.T) {
// 	fmt.Printf("domain: %s\noutput: %s\ndomains: %s\n", domain, output, domains)
// 	fmt.Println("AssetFinder")
// 	assets := fetchAssetFinder(domain)
// 	fmt.Println(assets)
// 	fmt.Println("AmassEnum")
// 	amassEnum := fetchAmassEnum(domain, output)
// 	fmt.Println(amassEnum)
// 	fmt.Println("SSLMate")
// 	sslMate := fetchSSLMate(domain)
// 	fmt.Println(sslMate)
// 	fmt.Println("Ping")
// 	res := activePing(domain, 3, 2)
// 	fmt.Printf("%t\n", res)
// 	fmt.Println("WebDomain")
// 	webDomains := fetchActive(domains)
// 	fmt.Println(webDomains)
// 	fmt.Println("Takeovers")
// 	potentialTakeovers := fetchSubjack(domains)
// 	fmt.Println(potentialTakeovers)
// 	fmt.Println("HTTProbe")
// 	httpURLS := fetchHTTProbe(domains)
// 	fmt.Println(httpURLS)
// 	fmt.Println("WaybackURLs")
// 	waybackurls := fetchWaybackURLs(domains)
// 	fmt.Println(waybackurls)
// 	fmt.Println("DNSInfo")
// 	dnsInfo := fetchDNSInfo(domain, "8.8.8.8:53")
// 	fmt.Println(dnsInfo)
// 	fmt.Println("Nmap")
// 	nmapInfo := fetchNmap([]string{domain})
// 	fmt.Println(nmapInfo)
// 	fmt.Println("Gowitness")
// 	fetchGoWitness(domains)
// 	fmt.Println("Whois")
// 	whois := fetchWhois(domain)
// 	fmt.Println(whois)
//
// }
