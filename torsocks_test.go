package torsocks

import (
	"fmt"
	"testing"
)

func TestBadConnect(t *testing.T) {
	fmt.Println("Testing to connect through 127.0.0.1:9050, to herp.derp:80.")
	tp, err := NewTorGate()
	if err != nil {
		fmt.Println("Maybe you are not running Tor?")
		fmt.Println(err.Error())
		t.Fail()
	}
	_, err = tp.DialTor("herp.derp:80")
	if err == nil {
		fmt.Println("We actually managed to connect to herp.derp. This is odd, and probably WRONG.\n")
		fmt.Println(err.Error())
		t.Fail()
	}
}

func TestConnect(t *testing.T) {
	onion := "facebookcorewwwi.onion:80"
	fmt.Println("Testing to connect through 127.0.0.1:9050, to " + onion)

	tp, err := NewTorGate()
	if err != nil {
		fmt.Println("Maybe you are not running Tor?")
		fmt.Println(err.Error())
		t.Fail()
		return
	}
	fmt.Printf("Dialing to " + onion + "\n")
	_, err = tp.DialTor(onion)
	if err != nil {
		fmt.Println("Maybe spore.mcx (" + onion + ") is not listening. Error was: " + err.Error())
		fmt.Println(err.Error())
		t.Fail()
		return
	}
	fmt.Println("Connected to .onion successfully!")
}
