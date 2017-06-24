# goosrelease [![Build Status](https://travis-ci.org/cobaugh/goosrelease.svg?branch=master)](https://travis-ci.org/cobaugh/goosrelease)

A Go package to make reading in os-release files easy.

See https://www.freedesktop.org/software/systemd/man/os-release.html

## Installation
`$ go get github.com/cobaugh/gooosrelease`

## Usage

```golang
package main

import (
	"fmt"
	"github.com/cobaugh/goosrelease"
)

func main() {
	// for reference, two variables are provided:
	fmt.Printf("EtcOsRelease = %v\n", goosrelease.EtcOsRelease)
	fmt.Printf("UsrLibOsRelease = %v\n", goosrelease.UsrLibOsRelease)

	// let goosrelease find what file to load
	osrelease, err := goosrelease.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("PRETTY_NAME = %v\n", osrelease["PRETTY_NAME"])

	// specify the file to load explicitly
	osrelease, err = goosrelease.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("PRETTY_NAME = %v\n", osrelease["PRETTY_NAME"])
}

```

Output:
```
$ ./examples 
EtcOsRelease = /etc/os-release
UsrLibOsRelease = /usr/lib/os-release
PRETTY_NAME = void
PRETTY_NAME = void```
