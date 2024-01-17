//
// Copyright 2014-2017 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package main

import (
	"bufio"
	"log"
	"os"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/smoltools/lsusbtty/usbID"
	"go.bug.st/serial.v1/enumerator"
)

func main() {
	idl, err := usbID.DownloadList("http://www.linux-usb.org/usb.ids")
	if err != nil {
		color256.PrintOrange("Error while downloading the id list")
	}
	color256.PrintBgHiMagenta("lsusbtty V1.1")

	color256.PrintHiCyan("Found USB-Serial converter")
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		color256.PrintRed("No serial ports found!")
		return
	}
	for _, port := range ports {
		if port.IsUSB {
			color256.PrintGreen("Connected adapter: %s", port.Name)
			color256.PrintWhite("   USB ID: %s:%s SN: %s", port.VID, port.PID, port.SerialNumber)
			if idl != "" {
				dev := usbID.GetID(port.VID, port.PID, idl)
				for _, d := range dev {
					if d != "" {
						color256.PrintWhite("     " + d)
					}
				}
			}

		} else {
			color256.PrintYellow("Not connected adapter: %s", port.Name)
		}
	}
	color256.PrintCyan("Press any key to close")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
