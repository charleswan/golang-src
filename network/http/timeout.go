package main

import (
	"net"
	"net/http"
	"time"
)

const httpRequestTimeout = 300 // 5 minute
const tcpConnectTimeout = 60   // 60 seconds
const tlsHandshakeTimeout = 60 // 60 seconds

// GetTimeoutNetClient ...
func GetTimeoutNetClient() *http.Client {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: tcpConnectTimeout * time.Second,
		}).Dial,
		TLSHandshakeTimeout: tlsHandshakeTimeout * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   httpRequestTimeout * time.Second,
		Transport: netTransport,
	}
	return netClient
}
