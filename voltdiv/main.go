package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
)

func main() {
	banner := `
Prints a table of minimal and maximal values of Uin and Uout of a voltage divider.

https://github.com/morgulbrut/goexperiments -> voltdiv

Uin ──────┐
          │
         ┌┴┐
         │ │ R1
         └┬┘
          │
          ├─────── Uout
          │
         ┌┴┐
         │ │ R2
         └┬┘
          │
   ───────┴───────
`

	parser := argparse.NewParser("voltdiv", banner)
	// Create string flag
	uins := parser.Float("i", "uin", &argparse.Options{Default: 0.0, Help: "Input voltage"})
	uouts := parser.Float("o", "uout", &argparse.Options{Default: 0.0, Help: "Output voltage"})
	r1s := parser.String("r", "r1", &argparse.Options{Required: true, Help: "Resistor R1, in the form 10, 1.2M"})
	r2s := parser.String("s", "r2", &argparse.Options{Required: true, Help: "Resistor R2, in the form 10, 1.2M"})
	tol := parser.Float("p", "percent", &argparse.Options{Default: 1.0, Help: "Tolerance for the resistors in percent"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// parse resistor values
	r1 := parseResistor(*r1s)
	r2 := parseResistor(*r2s)

	// renders table
	printVoltages(*uins, *uouts, r1, r2, *tol)
}

// printVoltages calculates and prints out the table
func printVoltages(ui, uo, r1, r2, tol float64) {

	// calculates minimal and maximal resistor values
	r1min := r1 - r1/100*tol
	r1max := r1 + r1/100*tol
	r2min := r2 - r2/100*tol
	r2max := r2 + r2/100*tol

	// calculates voltage ratios
	h := calcH(r1, r2)
	hmax := calcH(r1max, r2min)
	hmin := calcH(r1min, r2max)

	// calculates voltage values
	uIn, uOut := voltage(ui, uo, h)
	uInmin, uOutmin := voltage(ui, uo, hmin)
	uInmax, uOutmax := voltage(ui, uo, hmax)

	// sets up table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{fmt.Sprintf("Tol: %0.1f%%", tol), "R1", "R2", "H", "Uin", "Uout"})

	// fills table values
	t.AppendRow(table.Row{"Minimal", r1min, r2min, fmt.Sprintf("%0.5f", hmin), fmt.Sprintf("%0.3f", uInmin), fmt.Sprintf("%0.3f", uOutmin)})
	t.AppendRow(table.Row{"Nominal", r1, r2, fmt.Sprintf("%0.5f", h), fmt.Sprintf("%0.3f", uIn), fmt.Sprintf("%0.3f", uOut)})
	t.AppendRow(table.Row{"Maximal", r1max, r2max, fmt.Sprintf("%0.5f", hmax), fmt.Sprintf("%0.3f", uInmax), fmt.Sprintf("%0.3f", uOutmax)})

	t.Render()
}

// voltage calculates the Vin and Vout depending which voltage is given
func voltage(ui, uo, h float64) (uIn, uOut float64) {
	// calculates Vout
	if uo == 0.0 {
		uo = ui * h
	}
	// calculates Vin
	if ui == 0.0 {
		if h != 0.0 {
			ui = uo / h
		} else {
			color.Red("h is 0 => division by 0")
		}
	}
	return ui, uo
}

// parseResistor parses the resistor string into a float64
func parseResistor(r string) float64 {
	ret := 0.0

	if strings.HasSuffix(r, "G") {
		ret, _ = strconv.ParseFloat(strings.TrimRight(r, "G"), 64)
		ret *= 1e9
	} else if strings.HasSuffix(r, "M") {
		ret, _ = strconv.ParseFloat(strings.TrimRight(r, "M"), 64)
		ret *= 1e6
	} else if strings.HasSuffix(r, "k") {
		ret, _ = strconv.ParseFloat(strings.TrimRight(r, "k"), 64)
		ret *= 1e3
	} else {
		ret, _ = strconv.ParseFloat(r, 64)
	}
	return ret
}

// calcH calculates the voltage ration based on R1 and R2
func calcH(r1, r2 float64) float64 {
	return r2 / (r1 + r2)
}
