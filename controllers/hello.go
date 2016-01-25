package controllers

import (
	"github.com/martini-contrib/render"
	_ "github.com/martini-contrib/sessions"
	"net"
	"os"
)

func Hello(r render.Render) {
	ips := make([]string, 3)
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ips = append(ips, ipv4.String())
		}
	}
	r.JSON(200, ips)
}
