package converter

import (
	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/toml"
)

func PrintKeymap(km map[string]interface{}) {
	for k, val := range km {
		color256.PrintHiGreen("%s: %s", k, val)
	}
}

func ProcessKeymap(filename string) map[string]interface{} {
	var result map[string]interface{}
	// Unmarshal or Decode the JSON to the interface.
	val := readfile(filename)
	toml.Unmarshal(val, &result)

	return result
}
