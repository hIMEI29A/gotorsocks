# gotorsocks

**gotorsocks** - Easy to use Golang's **torify**. Provides CLI access to **Tor** proxying inside of **Go apps**. So, can parse web-sites on the 
**.onion domens**.

It is GitHub **fork** of [this BitBucket's repo](https://bitbucket.org/kallevedin/torsocks).

Old import path "code.google.com/p/go.net/proxy" (same as "golang.org/x/net/proxy") used in original version, is deadly broken, so original 
package uninstallable. In current version import path corrected, and some detail of code is changed.

[![Go Report Card](https://goreportcard.com/badge/github.com/hIMEI29A/gotorsocks)](https://goreportcard.com/report/github.com/hIMEI29A/gotorsocks) [![GoDoc](https://godoc.org/github.com/hIMEI29A/gotorsocks?status.svg)](http://godoc.org/github.com/hIMEI29A/gotorsocks)



## Install

Project uses [Glide](https://glide.sh) for dependencies manage. 
So, **install Glide first**:

```sh
curl https://glide.sh/get | sh
```
**Clone** repo:

```sh
git clone https://github.com/hIMEI29A/gotorsocks.git

cd gotorsocks
```

**Update** deps and install:

```sh
glide update

glide install
```
It will install dependencies ("golang.org/x/net/proxy" only) to `vendor/` folder of the repo.

**Run**

```sh
make
```

It will **run tests** and if it passed, install **gotorsocks** in to your $GOPATH

Also you can simply

```sh
go get github.com/hIMEI29A/gotorsocks
```

Check the [release page](https://github.com/hIMEI29A/gotorsocks/releases)!

## Example

```go
package main

import (
    "bufio"
    "fmt"

    "github.com/hIMEI29A/gotorsocks"
)

func main() {
    address := "facebookcorewwwi.onion:80"
    tor, err := gotorsocks.NewTorGate()

    if err != nil {
        panic(err)
    }

    connect, err := tor.DialTor(address)

    if err != nil {
        panic(err)
    }

    fmt.Fprintf(connect, "GET / HTTP/1.0\r\n\r\n")
    status, err := bufio.NewReader(connect).ReadString('\n')
    fmt.Println(status)
}
```
