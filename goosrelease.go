package goosrelease

import (
	"bufio"
	"errors"
	"os"
	"strings"
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
	err = nil

	lines, err := readFile(filename)
	if err != nil {
		return
	}

	for _, v := range lines {
		key, value, err := parseLine(v)
		if err == nil {
			osrelease[key] = value
		}
	}
	return
}

func parseLine(line string) (key string, value string, err error) {
	err = nil

	if len(line) == 0 {
		err = errors.New("Skipping: zero-length")
		return
	}
	if line[0] == '#' {
		err = errors.New("Skipping: comment")
		return
	}

	splitString := strings.SplitN(line, "=", 2)
	if len(splitString) != 2 {
		err = errors.New("Can not extract key=value")
		return
	}

	key = splitString[0]
	key = strings.Trim(key, " ")

	value = splitString[1]
	value = strings.Trim(value, " ")

	if strings.ContainsAny(value, `"`) {
		first := string(value[0:1])
		last := string(value[len(value)-1:])

		if first == last && strings.ContainsAny(first, `"'`) {
			value = strings.Trim(value, `"'`)
		}
		value = strings.Replace(value, `\"`, `"`, -1)
	}
	value = strings.Replace(value, `\$`, `$`, -1)
	value = strings.Replace(value, `\\`, `\`, -1)
	value = strings.Replace(value, "\\`", "`", -1)
	return
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
