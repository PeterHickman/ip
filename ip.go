package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"slices"
	"strings"
)

const address_service string = "http://ifconfig.me"

func local_addresses() string {
	// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go

	addresses := []string{}

	ifaces, _ := net.Interfaces()

	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if !ip.IsLoopback() && ip.IsPrivate() && ip.To4() != nil {
				addresses = append(addresses, ip.String())
			}
		}
	}

	if len(addresses) != 0 {
		// We sort this to ensure some sort of consistency
		slices.Sort(addresses)
		return addresses[0]
	} else {
		return "0.0.0.0"
	}
}

func public_address() string {
	res, err := http.Get(address_service)

	if err != nil {
		usage(err.Error())
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		usage(err.Error())
	}

	return string(resBody)
}

func usage(message string) {
	if message != "" {
		fmt.Println(message)
		fmt.Println()
	}

	fmt.Println("Find your local or public ip address")
	fmt.Println()
	fmt.Println("ip local - returns your local / private ip address")
	fmt.Println("ip public - returns the ip address that the internet can see")

	os.Exit(1)
}

func main() {
	flag.Parse()

	var address string

	if len(flag.Args()) == 1 {
		switch strings.ToLower(flag.Arg(0)) {
		case "local":
			address = local_addresses()
		case "public":
			address = public_address()
		default:
			usage("")
		}
	} else {
		usage("")
	}

	fmt.Print(address)
}
