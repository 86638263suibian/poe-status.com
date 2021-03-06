package scanner

import (
	"log"
	"strings"

	"github.com/Gonzih/poe-status.com/sh"
)

// NmapAvailable will test if nmap is present on the system
func NmapAvailable() bool {
	out, err := sh.Sh("bash", "-c", "which nmap")
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(strings.Trim(string(out), "\r\n "))
	return true
}

// PingAvailable will test if ping is present on the system
func PingAvailable() bool {
	out, err := sh.Sh("bash", "-c", "which ping")
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(strings.Trim(string(out), "\r\n "))
	return true
}
