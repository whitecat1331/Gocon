package main

import (
	"fmt"
	"os"

	"github.com/whitecat1331/amass/v4/amass"
	"github.com/whitecat1331/assetfinder"
	"github.com/whitecat1331/gowitness"
)

func setArgs(sys_args []string) []string {
	temp_args := os.Args
	os.Args = os.Args[0:1]
	os.Args = append(os.Args, sys_args...)
	return temp_args
}

func fetchAssetFinder(domain string) []string {
	return assetfinder.AssetFinder(domain, true)
}

func fetchAmassEnum(domain string, output string) []string {
	return amass.EnumAllDomain(domain, output)
}

func fetchGoWitness(urls string) {
	arg := "single"
	temp_args := setArgs([]string{arg})
	gowitness.GoWitness()
	os.Args = temp_args
}

func test() {
	domain := "zoom.com"
	output := "gocon.txt"
	// assets := fetchAssetFinder("youtube.com")
	amassEnum := fetchAmassEnum(domain, output)
	fmt.Println(amassEnum)
}

func main() {
	test()
}
