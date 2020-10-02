//
// Copyright 2014-2017 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package main

import (
	"log"
	"os"
	"bufio"

	"go.bug.st/serial.v1/enumerator"
	"github.com/fatih/color"

	"github.com/morgulbrut/usbIds"
)



func main() {	
	color.Cyan("Found USB-Serial converter")
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		color.Red("No serial ports found!")
		return
	}
	for _, port := range ports {
		if port.IsUSB {
			color.Green("Connected adapter: %s", port.Name)
			color.White("   USB ID: %s:%s SN: %s", port.VID, port.PID, port.SerialNumber)
			color.White(usbIds.GetId(port.VID,port.PID))
		} else {
			color.Yellow("Not connected adapter: %s",port.Name)
		}
	}
	color.Cyan("Press any key to close")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}