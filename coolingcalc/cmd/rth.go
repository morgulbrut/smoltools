package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type RthCmd struct {
	w     float64
	h     float64
	th    float64
	lamda float64
}

func (*RthCmd) Name() string { return "rth" }
func (*RthCmd) Synopsis() string {
	return "Calculates the thermal resistance from thermal conductivity"
}
func (*RthCmd) Usage() string { return "See \"flags rth\"\n\n" }

func (p *RthCmd) SetFlags(f *flag.FlagSet) {
	f.Float64Var(&p.w, "w", 100, "Width in mm")
	f.Float64Var(&p.h, "h", 100, "Height in mm")
	f.Float64Var(&p.th, "th", 1.6, "Thickness in mm")
	f.Float64Var(&p.lamda, "lambda", 2, "Thermal conductivity in K/m*W")
}

func (p *RthCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	l := p.th / 1000
	A := (p.w / 1000) * (p.h / 1000)
	fmt.Printf("Rth = %f\n", l/(p.lamda*A))
	return subcommands.ExitSuccess
}
