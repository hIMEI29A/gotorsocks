---
title: "gotorsocks - Easy to use Golang's torify"
date: "2017-12-01T22:13:12-05:00"
categories: ["Network", "Anonimity"]
tags: ["go","programming","gotorsocks","tor"]
---
# {{.Name}}

**gotorsocks** - Easy to use Golang's **torify**. Gets CLI access to **Tor** proxying inside of **Go apps**. So, parses web-sites on the 
**.onion domens**.

It is **GitHub-fork** of [BitBucket's repo](https://bitbucket.org/kallevedin/torsocks).

Old import path "code.google.com/p/go.net/proxy" (same as "golang.org/x/net/proxy") used in original version, is fatal broken, so original 
package uninstallable. In current version import path corrected, and some detail of code is changed.

{{template "badge/goreport" .}} {{template "badge/godoc" .}}

{{pkgdoc}}

## Install

Project uses [Glide](https://glide.sh) for dependencies manage. 
So, **install Glide first**:

```sh
curl https://glide.sh/get | sh
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

{{template "gh/releases" .}}

#### go
{{template "go/install" .}}
