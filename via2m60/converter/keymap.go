package converter

import (
	"fmt"
	"strings"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/toml"
)

func PrintKeymap(km map[string]interface{}) {
	for k, val := range km {
		color256.PrintHiGreen("%s: %s", k, val)
	}
}

func ReadKeymap(filename string) map[string]interface{} {
	var result map[string]interface{}
	// Unmarshal or Decode the JSON to the interface.
	val := readfile(filename)
	toml.Unmarshal(val, &result)

	return result
}

func parseTabs(key string) (tapped string, hold string) {
	a := strings.Split(key, "(")[1]
	a = strings.Split(a, ")")[0]
	t := strings.Split(a, ",")[0]
	h := strings.Split(a, ",")[1]
	return t, h
}

func GenerateLayers(config ViaJSON, km map[string]interface{}) {
	var lays strings.Builder
	for _, l := range config.Layers {
		for i, k := range l {
			if strings.HasPrefix(k, "MT") {
				t, h := parseTabs(k)
				lays.WriteString(fmt.Sprintf("MODS_TAP(MODS(%s,%s)", km[t], km[h]))
			} else if strings.HasPrefix(k, "LT") {
				t, h := parseTabs(k)
				lays.WriteString(fmt.Sprintf("LAYER_TAB(%s,%s)", t, km[h]))
			} else if km[k] != "NO" {
				lays.WriteString(fmt.Sprintf("%s,", km[k]))
			}
			// TODO:
			if i == 14 || i == 29 || i == 43 || i == 58 {
				lays.WriteString("\n")
			}
		}
		lays.WriteString("\n--------------------------------------\n")
	}
	fmt.Println(lays.String())
}
