package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	ip_str := ""
	for _, b := range ip {
		ip_str += fmt.Sprint(b) + "."
	}
	return ip_str[:len(ip_str)-1]
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
