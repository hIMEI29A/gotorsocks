// Here is is GitHub fork of https://bitbucket.org/kallevedin/torsocks.
// Import path "code.google.com/p/go.net/proxy" (same as "golang.org/x/net/proxy")
// used in original version, is fatal broken, so packege uninstallabnle.
// In current version import path corrected, and some detail of code is changed
//
// https://bitbucket.org/kallevedin/torsocks is relised on Public Domain.

package torsocks

import (
	"bytes"
	"errors"
	"net"
	"time"

	"golang.org/x/net/proxy"
)

// A Tor proxy. Is actually just a string with the address of the Tor Proxy. (Needs to be an IPv4
// address or a domain name that can be translated to an IPv4 address, with a port.)
// Examples: "127.0.0.1:9050", "10.0.30.11:9150".
type TorProxy string

// Defines a new Tor SOCKS5 proxy. Also performs very basic heuristic check to see if it actually
// is a Tor proxy we are connecting to, and not a normal SOCKS proxy, or something else.
func NewTorProxy(address string) (*TorProxy, error) {
	duration, _ := time.ParseDuration("10s")
	conn, err := net.DialTimeout("tcp4", address, duration)
	if err != nil {
		return nil, errors.New("Could not test Tor Proxy: " + err.Error())
	}
	// Tor proxies reply to anything that looks like HTTP GET or POST with known error message.
	conn.Write([]byte("GET /\n"))
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return nil, errors.New("Failed heuristics. " + address + " is maybe not a Tor Proxy.")
		}
		if bytes.Contains(buf[:n], []byte("Tor is not an HTTP Proxy")) {
			conn.Close()
			tp := TorProxy(address)
			return &tp, nil
		}
	}
}

// Defines a new Tor SOCKS5 proxy. Does not perform any heuristic checks, and always succeeds.
func NewTorProxyNoCheck(address string) *TorProxy {
	t := TorProxy(address)
	return &t
}

// Dials to the address, through the Tor Proxy. address needs to be an IPv4 address or a domain name
// that can be translated to an IPv4 address (or a .onion address!) followed by the port number you
// want to connect to. Examples: "1.2.3.4:22", "duckduckgo.com:80", "227vftpsbp62v7bd.onion:6667".
func (tp *TorProxy) DialTor(address string) (net.Conn, error) {
	dialer, err := proxy.SOCKS5("tcp4", string(*tp), nil, proxy.Direct)
	if err != nil {
		return nil, errors.New("Could not connect to Tor proxy SOCKS: " + err.Error())
	}
	conn, err := dialer.Dial("tcp4", address)
	if err != nil {
		return nil, errors.New("Failed to connect: " + err.Error())
	}
	return conn, nil
}
