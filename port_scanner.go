package main

import (
	"net"
	"strconv"
	"time"
)

func scanPorts(subdomain Subdomain) Subdomain {
	socketAddresses, err := net.LookupHost(subdomain.domain + ":1024")
	if err != nil {
		return subdomain
	}

	var openPorts []Port
	for _, address := range socketAddresses {
		for _, port := range MostCommonPorts100 {
			openPorts = append(openPorts, scanPort(address, port))
		}
	}
	subdomain.openPorts = openPorts
	return subdomain
}

func scanPort(address string, port uint16) Port {
	timeout := time.Duration(3 * time.Second)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(address, strconv.Itoa(int(port))), timeout)
	isOpen := err == nil
	if conn != nil {
		conn.Close()
	}
	return Port{port: port, isOpen: isOpen}
}
