package gocon

import (
	"github.com/whitecat1331/waybackurls"
)

var waybackurlsLog = setLogName("waybackurls.log")

func fetchWaybackURLs(domains []string, logPath string) func() ([]string, error) {
	return func() ([]string, error) {
		return waybackurls.WaybackURLS(domains, logPath)
	}
}

func FetchWaybackURLs(domains []string) ([]string, error) {
	return fetchWaybackURLs(domains, waybackurlsLog)()
}
