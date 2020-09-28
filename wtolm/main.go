package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/morgulbrut/smoltools/wtolm/lut"
)

/*
θv = Photopic Luminous Flux (lm)
Km = 683 lm/W
V(λ) = the "luminous efficiency"
θλ = Spectral Radiant Flux (W)

θv = θλ * Km * V(λ)

*/

const Km float64 = 685

func main() {
	var θλ float64
	var λ int
	flag.Float64Var(&θλ, "p", 0.1, "θλ, Spectral Radiant Flux in mW")
	flag.IntVar(&λ, "l", 550, "λ Wavelenght in nm rounded to the next")

	// custom help string
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "%s\n", os.Args[0])
		fmt.Fprintf(os.Stdout, "Smol tool to calculate lm from Watts.\n\n")
		fmt.Fprintf(os.Stdout, "Usage:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Println(θλ * Km * lut.Vλ[λ])
}
