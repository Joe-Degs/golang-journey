package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

/*
   An IP address is typically divided into the components of a network address, a subnet and a device
   portion.
   The network address and subnet form a prefix to the device portion.
   The mask is an IP address of all binary ones to match the prefix length, followed by all zeros.
   For an IPv4 address of 103.232.159.187 on a /24 network, this means that the mask is made up of
   24 ones followed by 8 zeros because the IP addr is 32bits long.
*/

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr ones bits\n", os.Args[0])
		os.Exit(-1)
	}
	dotAddr := os.Args[1]
	ones, _ := strconv.Atoi(os.Args[2])
	bits, _ := strconv.Atoi(os.Args[3])

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid Address")
		os.Exit(1)
	}
	mask := net.CIDRMask(ones, bits)
	network := addr.Mask(mask)
	//calculated usable ip range.
	fmt.Println("Address is ", addr.String(),
		"\nMask length is ", bits,
		"\nLeading ones count is ", ones,
		"\nMask is (hex) ", mask.String(),
		"\nNetwork is ", network.String(),
		"\nThe number of usable IPs is ", usableIP(ones),
	)
	os.Exit(0)
}

func usableIP(ones int) int {
	zeros := 32 - ones
	result := 2
	for i := 1; i < zeros; i++ {
		result *= 2
	}
	return result - 2
}
