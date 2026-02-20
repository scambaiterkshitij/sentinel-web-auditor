package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

type Result struct {
	Target             string        `json:"target"`
	IP                 string        `json:"ip"`
	PortOpen           bool          `json:"port_open"`
	TLSVersion         string        `json:"tls_version"`
	CertificateIssuer  string        `json:"certificate_issuer"`
	CertificateExpiry  string        `json:"certificate_expiry"`
	HTTPStatus         int           `json:"http_status"`
	Headers            http.Header   `json:"headers"`
	Latency            time.Duration `json:"latency_ms"`
	Error              string        `json:"error,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: netprobe <target>")
		return
	}

	target := os.Args[1]
	result := Result{Target: target}

	start := time.Now()

	ips, err := net.LookupIP(target)
	if err != nil || len(ips) == 0 {
		result.Error = "DNS resolution failed"
		output(result)
		return
	}
	result.IP = ips[0].String()

	conn, err := net.DialTimeout("tcp", target+":443", 5*time.Second)
	if err != nil {
		result.PortOpen = false
		result.Error = "Port 443 closed"
		output(result)
		return
	}
	result.PortOpen = true
	conn.Close()

	conf := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         target,
	}

	tlsConn, err := tls.Dial("tcp", target+":443", conf)
	if err == nil {
		state := tlsConn.ConnectionState()
		result.TLSVersion = tlsVersion(state.Version)

		if len(state.PeerCertificates) > 0 {
			cert := state.PeerCertificates[0]
			result.CertificateIssuer = cert.Issuer.CommonName
			result.CertificateExpiry = cert.NotAfter.String()
		}
		tlsConn.Close()
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://" + target)
	if err == nil {
		result.HTTPStatus = resp.StatusCode
		result.Headers = resp.Header
		resp.Body.Close()
	}

	result.Latency = time.Since(start)

	output(result)
}

func tlsVersion(v uint16) string {
	switch v {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown"
	}
}

func output(r Result) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(r)
}
