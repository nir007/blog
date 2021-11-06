package utils

import (
	"math/rand"
	"net"
	"regexp"
	"strconv"
	"time"
)

// GetConformationCode generates random code string
func GetConformationCode() string {
	rand.Seed(time.Now().UTC().UnixNano())
	code := strconv.Itoa(rand.Intn(100)) + strconv.Itoa(rand.Intn(1000))
	return code
}

// GetIP returns client ip address
func GetIP() string {
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

// ThoroughlyClearString clear string
func ThoroughlyClearString(str string) (res string) {
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	res = re.ReplaceAllString(str, "")
	return res
}
