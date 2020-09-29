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
	var θv float64
	var λ int
	flag.Float64Var(&θλ, "p", 0, "θλ, Spectral Radiant Flux in mW")
	flag.Float64Var(&θv, "l", 0, "θv, Photopic Luminous Flux in lm")
	flag.IntVar(&λ, "w", 550, "λ Wavelenght in nm rounded to the next")

	// custom help string
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "%s\n\n", os.Args[0])
		fmt.Fprintf(os.Stdout, "Smol tool to calculate lumens from watts.\nAnd vice versa.\n\n")
		fmt.Fprintf(os.Stdout, "Usage:\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if θλ == 0 && θv == 0 {
		flag.Usage()
	} else if θv == 0 {
		fmt.Println(θλ * Km * lut.Vλ[λ])
	} else if θλ == 0 {
		fmt.Println(θv / (Km * lut.Vλ[λ]))
	} else {
		flag.Usage()
	}
}
