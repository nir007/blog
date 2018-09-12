package helpers

import (
	"math/rand"
	"time"
	"strconv"
	"net"
)

type Helper struct {}

func (h *Helper) GetConformationCode() string {
	rand.Seed(time.Now().UTC().UnixNano())
	code := strconv.Itoa(rand.Intn(100))
	code += strconv.Itoa(rand.Intn(1000))
	return code
}

func (h * Helper) GetIp() string {
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				return v.IP.String()
			case *net.IPAddr:
				return v.IP.String()
			}
		}
	}
	return ""
}