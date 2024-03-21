package gocon

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type SSLMateJSON []struct {
	ID           string   `json:"id"`
	TbsSha256    string   `json:"tbs_sha256"`
	CertSha256   string   `json:"cert_sha256"`
	DNSNames     []string `json:"dns_names"`
	PubkeySha256 string   `json:"pubkey_sha256"`
	Issuer       struct {
		FriendlyName string `json:"friendly_name"`
		PubkeySha256 string `json:"pubkey_sha256"`
		Name         string `json:"name"`
	} `json:"issuer"`
	NotBefore  time.Time `json:"not_before"`
	NotAfter   time.Time `json:"not_after"`
	Revoked    bool      `json:"revoked"`
	Revocation struct {
		Time      any       `json:"time"`
		Reason    any       `json:"reason"`
		CheckedAt time.Time `json:"checked_at"`
	} `json:"revocation"`
	ProblemReporting string `json:"problem_reporting"`
	CertDer          string `json:"cert_der"`
}

func fetchSSLMate(domain string) []string {
	var dnsNames []string
	baseURL := "https://api.certspotter.com/v1/issuances?domain=%s&expand=dns_names&expand=issuer&expand=revocation&expand=problem_reporting&expand=cert_der"
	armedURL := fmt.Sprintf(baseURL, domain)

	resp, err := http.Get(armedURL)
	if err != nil {
		defer resp.Body.Close()
		log.Println(err)
		return dnsNames
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return dnsNames
	}

	var sslJSON SSLMateJSON

	err = json.Unmarshal(body, &sslJSON)
	if err != nil {
		log.Fatal(err)
		return dnsNames
	}

	for _, sslResponse := range sslJSON {
		dnsNames = append(dnsNames, sslResponse.DNSNames...)
		return dnsNames
	}

	return dnsNames
}
