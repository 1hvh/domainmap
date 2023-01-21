package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net"
	"strconv"
	"strings"
	"sync"
)

func enumerate(httpClient *resty.Client, target string) ([]Subdomain, error) {
	var entries []Cert
	var uri string = fmt.Sprintf("https://crt.sh/?q=%s&output=json", target)
	_, err := httpClient.R().SetResult(&entries).Get(uri)
	if err != nil {
		return nil, err
	}

	var subdomains []string
	for _, entry := range entries {
		for _, subdomain := range strings.Split(entry.NameValue, "\n") {
			subdomain = strings.TrimSpace(subdomain)
			if subdomain != target && !strings.Contains(subdomain, "*") && !contains(subdomains, subdomain) {
				subdomains = append(subdomains, subdomain)
			}
		}
	}

	var result []Subdomain = checkAvailability(subdomains)

	return result, nil
}

func checkAvailability(domains []string) []Subdomain {
	var wg sync.WaitGroup
	var subdomains []Subdomain

	for _, domain := range domains {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			var openPorts []Port
			fmt.Println("Checking " + domain + " for availability...")
			_, error := net.Dial("tcp", fmt.Sprintf("%s:%s", domain, "http"))
			if error == nil {
				fmt.Println("Scanning " + domain + " ports...")
				for _, port := range MostCommonPorts100 {
					conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", domain, "http"))
					if err == nil {
						openPorts = append(openPorts, Port{port: uint16(port), isOpen: true})
					}
					conn.Close()
				}
			}
			fmt.Println("Found " + strconv.Itoa(len(openPorts)) + " opened ports for " + domain)
			subdomains = append(subdomains, Subdomain{domain: domain, openPorts: openPorts})

		}(domain)
	}

	wg.Wait()
	return subdomains
}
