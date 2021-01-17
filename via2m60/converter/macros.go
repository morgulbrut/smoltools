package converter

import (
	"strings"
)

func GenerateMacros(config ViaJSON) [][]string {
	var macros [][]string
	for _, m := range config.Macros {
		var macro []string
		if m != "" {
			macro = parseMacro(m)
		}
		macros = append(macros, macro)
	}
	return macros
}

func parseMacro(m string) []string {
	var ret []string
	m = strings.ReplaceAll(m, "}", "")
	ret = strings.Split(m, "{")
	return ret
}
