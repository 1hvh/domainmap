package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/panjf2000/ants"
	"os"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: <program> <target>")
		return
	}
	target := args[1]

	httpTimeout := time.Duration(5 * time.Second)
	httpClient := resty.New().SetTimeout(httpTimeout).SetRedirectPolicy(resty.FlexibleRedirectPolicy(4))

	p, _ := ants.NewPool(256)
	defer p.Release()

	var scanResult []Subdomain
	scanResult, _ = enumerate(httpClient, target)
	fmt.Println("---------------------------")
	fmt.Println("Results:")
	fmt.Println("---------------------------")
	for _, subdomain := range scanResult {
		fmt.Printf("%s ", subdomain.domain)
		for _, port := range subdomain.openPorts {
			fmt.Printf("%d%s", port.port, ",")
		}
		fmt.Printf("%s", "\n\n")
	}
}
