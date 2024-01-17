package main

import (
	"os"

	"github.com/beevik/ntp"
	"github.com/morgulbrut/color256"
)

func main() {
	// Specify the NTP server address
	ntpServer := "pool.ntp.org"
	if len(os.Args) > 1 {
		ntpServer = os.Args[1]
	}
	color256.PrintHiCyan("NTP Server: %s", ntpServer)

	// Get the current time from the NTP server
	ntpTime, err := ntp.Query(ntpServer)
	if err != nil {
		color256.PrintHiRed(err.Error())
	} else {
		color256.PrintHiGreen("NTP Time: %s", ntpTime.Time)
	}
}
