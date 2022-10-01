package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/joho/godotenv"
	"net"
	"os"
)

func IpIsVpn(ip string) bool {
	err := godotenv.Load()
	if err != nil {
		return false
	}
	client := ipinfo.NewClient(nil, nil, os.Getenv("IPINFO_TOKEN"))
	info, _ := client.GetIPPrivacy(net.ParseIP(ip))
	if info.VPN == true {
		return true
	}
	return false
}

func IpIsTor(ip string) bool {
	err := godotenv.Load()
	if err != nil {
		return false
	}
	client := ipinfo.NewClient(nil, nil, os.Getenv("IPINFO_TOKEN"))
	info, _ := client.GetIPPrivacy(net.ParseIP(ip))
	if info.Tor == true {
		return true
	}
	return false
}

func IpIsProxy(ip string) bool {
	err := godotenv.Load()
	if err != nil {
		return false
	}
	client := ipinfo.NewClient(nil, nil, os.Getenv("IPINFO_TOKEN"))
	info, _ := client.GetIPPrivacy(net.ParseIP(ip))
	if info.Proxy == true {
		return true
	}
	return false
}

func IpIsHosting(ip string) bool {
	err := godotenv.Load()
	if err != nil {
		return false
	}
	client := ipinfo.NewClient(nil, nil, os.Getenv("IPINFO_TOKEN"))
	info, _ := client.GetIPPrivacy(net.ParseIP(ip))
	if info.Hosting == true {
		return true
	}
	return false
}

func IpIsRelay(ip string) bool {
	err := godotenv.Load()
	if err != nil {
		return false
	}
	client := ipinfo.NewClient(nil, nil, os.Getenv("IPINFO_TOKEN"))
	info, _ := client.GetIPPrivacy(net.ParseIP(ip))
	if info.Relay == true {
		return true
	}
	return false
}

func IPCheck(req *gin.Context) {
	ip := req.Param("ip")
	req.JSON(200, gin.H{
		"ip":      ip,
		"vpn":     IpIsVpn(ip),
		"proxy":   IpIsProxy(ip),
		"tor":     IpIsTor(ip),
		"hosting": IpIsHosting(ip),
		"relay":   IpIsRelay(ip),
	})
}
