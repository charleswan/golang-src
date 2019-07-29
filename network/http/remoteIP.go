package main

import (
	"net"
	"net/http"
)

const (
	// XForwardedFor ...
	XForwardedFor = "X-Forwarded-For"
	// XRealIP ...
	XRealIP = "X-Real-IP"
)

// RemoteIP 返回远程客户端的 IP，如 192.168.1.1
func RemoteIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

func main() {
}
