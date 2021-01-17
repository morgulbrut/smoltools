package main

import (
	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/smoltools/via2m60/converter"
)

func logo() {
	color256.Init()
	color256.PrintRandom("██╗   ██╗██╗ █████╗ ██████╗ ███╗   ███╗ ██████╗  ██████╗ ")
	color256.PrintRandom("██║   ██║██║██╔══██╗╚════██╗████╗ ████║██╔════╝ ██╔═████╗")
	color256.PrintRandom("██║   ██║██║███████║ █████╔╝██╔████╔██║███████╗ ██║██╔██║")
	color256.PrintRandom("╚██╗ ██╔╝██║██╔══██║██╔═══╝ ██║╚██╔╝██║██╔═══██╗████╔╝██║")
	color256.PrintRandom(" ╚████╔╝ ██║██║  ██║███████╗██║ ╚═╝ ██║╚██████╔╝╚██████╔╝")
	color256.PrintRandom("  ╚═══╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝  ╚═════╝ ")
	//color256.PrintRandom("  Turn your CircuitPython board into a USB RubberDucky")
}

func main() {
	logo()
	var o converter.Output
	config := converter.ReadJSON("dz60_via.json")
	converter.PrintConfig(config)
	km := converter.ReadKeymap("keycodes.toml")
	layers := converter.GenerateLayers(config, km)
	o.Keymap = converter.FmtLayers(layers)
	macros := converter.GenerateMacros(config)
	o.Macros = converter.FmtMacros(macros, km)
	converter.ExecuteTemplate(o, "code_template.py")
}
