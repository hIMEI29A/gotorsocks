// @Package gotorsocks
// @Author: Kalle Vedin <kalle.vedin@fripost.org>
// @Author: hIMEI <himei@tuta.io>
// @Date:   2017-12-16 22:02:59
// @Copyright © 2017 hIMEI <himei@tuta.io>
// @license MIT

// gotorsocks Here is GitHub fork of https://bitbucket.org/kallevedin/torsocks.
// Import path "code.google.com/p/go.net/proxy" (same as "golang.org/x/net/proxy")
// used in original version, is fatal broken, so original package uninstallable.
// In current version import path corrected, and some detail of code is changed
//
// https://bitbucket.org/kallevedin/torsocks is relised on Public Domain.

package gotorsocks

import (
	"bytes"
	"errors"
	"net"
	"time"

	"golang.org/x/net/proxy"
)

// TorGate is  a Tor proxy. Is actually just a string with the address of the Tor Proxy.
// (Needs to be an IPv4 address or a domain name that can be translated to an IPv4
// address, with a port.)
// Examples: "127.0.0.1:9050", "10.0.30.11:9150".
type TorGate string

// TOR_GATE string constant with localhost's Tor port
const TOR_GATE_ string = "127.0.0.1:9050"

// NewTorGate creates new TorGate
func NewTorGate() (*TorGate, error) {
	duration, _ := time.ParseDuration("10s")
	connect, err := net.DialTimeout("tcp4", TOR_GATE_, duration)

	if err != nil {
		return nil, errors.New("Could not test TOR_GATE_: " + err.Error())
	}

	// Tor proxies reply to anything that looks like
	// HTTP GET or POST with known error message.
	connect.Write([]byte("GET /\n"))
	connect.SetReadDeadline(time.Now().Add(10 * time.Second))
	buf := make([]byte, 4096)

	for {
		n, err := connect.Read(buf)

		if err != nil {
			return nil, errors.New("It is not TOR_GATE_")
		}

		if bytes.Contains(buf[:n], []byte("Tor is not an HTTP Proxy")) {
			connect.Close()
			gate := TorGate(TOR_GATE_)

			return &gate, nil
		}
	}
}

// DialTor dials to the .onion address
func (gate *TorGate) DialTor(address string) (net.Conn, error) {
	dialer, err := proxy.SOCKS5("tcp4", string(*gate), nil, proxy.Direct)

	if err != nil {
		return nil, errors.New("Could not connect to TOR_GATE_: " + err.Error())
	}

	connect, err := dialer.Dial("tcp4", address)

	if err != nil {
		return nil, errors.New("Failed to connect: " + err.Error())
	}

	return connect, nil
}
