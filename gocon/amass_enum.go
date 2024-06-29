package gocon

import (
	"github.com/whitecat1331/amass/v4/amass"
)

var amassLog = setLogName("amass_enum.log")

func fetchAmassEnum(cliArgs []string) func() []string {
	return func() []string {
		return amass.RunEnumCommand(cliArgs)
	}
}

func FetchAmassEnum(domains []string) []string {
	var fDomains string
	const minutes = 5
	for _, domain := range domains {
		fDomains += (domain + ",")
	}
	cliArgs := []string{"-silent", "-nocolor", "-timeout", string(minutes), "-dir", amassLog, "-d", fDomains}
	return fetchAmassEnum(cliArgs)()
}
