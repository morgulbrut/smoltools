package converter

import (
	"fmt"
	"strings"
)

func FmtLayers(lays [][]string) string {
	var ret strings.Builder
	ret.WriteString("keyboard.keymap = (\n")
	for i, l := range lays {
		ret.WriteString(fmt.Sprintf("    # layer %d\n    (\n        ", i))
		for j, k := range l {
			if j == 15 || j == 30 || j == 44 || j == 59 {
				ret.WriteString("\n        ")

			}
			if k == "TRANSPARENT" {
				ret.WriteString("___, ")
			} else if k != "" {
				ret.WriteString(fmt.Sprintf("%s, ", k))
			}
		}
		ret.WriteString("\n    ),\n")
	}
	ret.WriteString(")\n")
	return ret.String()
}
