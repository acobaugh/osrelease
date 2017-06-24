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
