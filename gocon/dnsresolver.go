package gocon

import (
	"context"
	"log"
	"net"
	"time"
)

// A AAAA MX NS TXT SOA

type DNSResolverInfo struct {
	A    []net.IP
	AAAA []net.IP
	MX   []*net.MX
	NS   []*net.NS
}

func NewDNSResolverInfo(A []net.IP, AAAA []net.IP, MX []*net.MX, NS []*net.NS) *DNSResolverInfo {
	return &DNSResolverInfo{
		A:    A,
		AAAA: AAAA,
		MX:   MX,
		NS:   NS,
	}

}

func fetchDNSInfo(domain string, dns string) *DNSResolverInfo {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network string, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, dns)
		},
	}

	ip4, err := r.LookupIP(context.Background(), "ip4", domain)
	if err != nil {
		log.Println(err)
	}

	ip6, err := r.LookupIP(context.Background(), "ip6", domain)
	if err != nil {
		log.Println(err)
	}

	mx, err := r.LookupMX(context.Background(), domain)

	if err != nil {
		log.Println(err)
	}

	ns, err := r.LookupNS(context.Background(), domain)

	if err != nil {
		log.Println(err)

	}

	return NewDNSResolverInfo(ip4, ip6, mx, ns)

}
