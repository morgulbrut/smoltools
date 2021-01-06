package converter

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/morgulbrut/color256"
)

// ViaJSON represents a VIA JSON file.
// At the moment we only need the the macros
// and the keymap
type ViaJSON struct {
	Name   string     `json:"name"`
	Macros []string   `json:"macros"`
	Layers [][]string `json:"layers"`
}

func PrintConfig(config ViaJSON) {
	color256.PrintHiPink("Name:      %s", config.Name)
	color256.PrintHiCyan("Layers:    %d", len(config.Layers))
}

func ProcessJSON(filename string) ViaJSON {
	var config ViaJSON

	file, err := os.Open(filename)
	if err != nil {
		color256.PrintHiRed(err.Error())
	}
	defer file.Close()
	val, err := ioutil.ReadAll(file)
	if err != nil {
		color256.PrintHiRed(err.Error())
	}
	json.Unmarshal(val, &config)

	return config
}
