---
title: "gotorsocks - Easy to use Golang's torify"
date: "2017-12-01T22:13:12-05:00"
categories: ["Network", "Anonimity"]
tags: ["go","programming","gotorsocks","tor"]
---
# {{.Name}}

**gotorsocks** - Easy to use Golang's **torify**. Gets CLI access to **Tor** proxying inside of **Go apps**. So, parses web-sites on the 
**.onion domens**.

It is GitHub **fork** of [this BitBucket's repo](https://bitbucket.org/kallevedin/torsocks).

Old import path "code.google.com/p/go.net/proxy" (same as "golang.org/x/net/proxy") used in original version, is deadly broken, so original 
package uninstallable. In current version import path corrected, and some detail of code is changed.

{{template "badge/goreport" .}} {{template "badge/godoc" .}}

{{pkgdoc}}

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

{{template "gh/releases" .}}

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
