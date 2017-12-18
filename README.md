# ipcheck
[![Build Status](https://img.shields.io/travis/mushroomsir/ipcheck.svg?style=flat-square)](https://travis-ci.org/mushroomsir/ipcheck)
[![Coverage Status](http://img.shields.io/coveralls/mushroomsir/ipcheck.svg?style=flat-square)](https://coveralls.io/github/mushroomsir/ipcheck?branch=master)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/mushroomsir/ipcheck/blob/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mushroomsir/ipcheck)

This repository lets you check if an IP matches one or more IP's or [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) ranges. It handles IPv6, IPv4, and IPv4-mapped over IPv6 addresses. 

## Features

- Support a single CIDR or IP string, e.g. "125.19.23.0/24", or "2001:cdba::3257:9652", or "62.230.58.1"
- Support an array of CIDR and/or IP strings, e.g. ["125.19.23.0/24", "2001:cdba::3257:9652", "62.230.58.1"]
- Custom file `Content-Type` for upload
- Support add extra field to formdata
- Download file 

## Installation

```sh
go get github.com/mushroomsir/ipcheck
```

## Usage
```go
package main

import "github.com/mushroomsir/ipcheck"

func main() {
	 ipcheck.IsRange("::1", "::2/128")
	 ipcheck.IsRange("2001:cdba::3257:9652", "2001:cdba::3257:9652/128")
}

```

## Licenses

All source code is licensed under the [MIT License](https://github.com/mushroomsir/ipcheck/blob/master/LICENSE).
