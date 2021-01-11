package converter

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/morgulbrut/color256"
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

func executeTemplate(s, temp string) {
	f, err := os.Create("code_via.py")
	if err != nil {
		color256.PrintHiRed("ERROR: template parsing fail")
		os.Exit(1)
	}
	t, err := template.ParseFiles(temp)
	if err != nil {
		color256.PrintHiRed("ERROR: template parsing fail")
		os.Exit(1)
	}
	t.Execute(f, s)
}
