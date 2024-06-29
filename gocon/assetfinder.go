package gocon

import (
	"github.com/whitecat1331/assetfinder"
)

var assetfinderLog = setLogName("assetfinder.log")

func fetchAssetFinder(domains []string, subsOnly bool, logPath string) func() ([]string, error) {
	return func() ([]string, error) {
		return assetfinder.AssetFinder(domains, subsOnly, logPath)
	}
}

func FetchAssetFinder(domains []string) ([]string, error) {
	return fetchAssetFinder(domains, true, assetfinderLog)()
}
