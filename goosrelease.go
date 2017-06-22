package goosrelease

import (
	"bufio"
	"fmt"
	"os"
)

const EtcOsRelease string = "/etc/os-release"

type OsRelease struct {
	NAME             string
	VERSION          string
	ID               string
	ID_LIKE          string
	VERSION_CODENAME string
	VERSION_ID       string
	PRETTY_NAME      string
	ANSI_COLOR       string
	CPE_NAME         string
	BUILD_ID         string
	VARIANT          string
	VARIANT_ID       string
}

func Read(filename string) (osrelease map[string]string, err error) {
	osrelease = make(map[string]string)

	lines, err := readFile(filename)
	for _, v := range lines {
		fmt.Print("%s\n", v)
	}
	return osrelease, nil
}

func readFile(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
