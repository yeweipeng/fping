package main

import (
	"net"
	"fmt"
)

func isIPv4 (ip string) bool {
	ipv4 := net.ParseIP(ip)
	fmt.Println(ipv4)
	return ipv4.To4() != nil
}

func main() {
	isIPv4("10.1.1.1")
	isIPv4("10.1.1.255")
	isIPv4("10.1.1.256")
	isIPv4("0.0.0.0")
}