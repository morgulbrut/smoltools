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

func GenerateLayers(config ViaJSON, km map[string]interface{}) [][]string {
	var lays [][]string
	for _, l := range config.Layers {
		var lay []string
		for _, k := range l {
			var key string
			if strings.HasPrefix(k, "MT") {
				t, h := parseTabs(k)
				key = fmt.Sprintf("MODS_TAP(MODS(%s),%s)", km[t], km[h])
			} else if strings.HasPrefix(k, "LT") {
				t, h := parseTabs(k)
				key = fmt.Sprintf("LAYER_TAP(%s,%s)", t, km[h])
			} else if km[k] != "NO" {
				if km[k] == nil {
					key = "NO"
				} else {
					key = fmt.Sprintf("%s", km[k])
				}
			}
			lay = append(lay, key)
		}
		lays = append(lays, lay)
	}
	return lays
}
