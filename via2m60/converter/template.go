package converter

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/morgulbrut/color256"
)

type Output struct {
	Keymap string
	Macros string
}

func find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func FmtMacros(macros [][]string, km map[string]interface{}) string {
	var ret strings.Builder
	// ret.WriteString("def macro_handler(dev, n, is_down):\n")
	// ret.WriteString("    if is_down:\n")
	// ret.WriteString("        pass\n")
	// ret.WriteString("    else:\n")
	for i, m := range macros {
		if i == 0 {
			ret.WriteString("        if n == 0:\n")
		} else {
			ret.WriteString(fmt.Sprintf("        elif n == %d:\n", i))
		}
		if m == nil {
			ret.WriteString("            pass\n")
		} else {
			for _, k := range m {
				if strings.HasPrefix(k, "KC_") {
					ret.WriteString(fmt.Sprintf("            dev.send("))
					keys := strings.Split(k, ",")
					for j, key := range keys {
						if j < len(keys)-1 {
							ret.WriteString(fmt.Sprintf("%s, ", km[key]))
						} else {
							ret.WriteString(fmt.Sprintf("%s", km[key]))
						}
					}
					ret.WriteString(")\n")
				} else {
					if k != "" {
						ret.WriteString(fmt.Sprintf("            dev.send_text(\"%s\")\n", k))
					}
				}
			}
		}
	}
	return ret.String()
}

func FmtLayers(lays [][]string) string {
	var ret strings.Builder
	// TODO:
	neededKeys := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14,
		15, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
		45, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 58,
		60, 61, 63, 66, 70, 71, 73, 74,
	}
	lineBreaks := []int{
		15, 30, 45, 60,
	}

	ret.WriteString("keyboard.keymap = (\n")
	for i, l := range lays {
		ret.WriteString(fmt.Sprintf("    # layer %d\n    (\n        ", i))
		for j, k := range l {
			if find(lineBreaks, j) {
				ret.WriteString("\n        ")
			}
			if find(neededKeys, j) {
				if k == "TRANSPARENT" {
					ret.WriteString("___, ")
				} else if k != "" {
					ret.WriteString(fmt.Sprintf("%s, ", k))
				}
			}
		}
		ret.WriteString("\n    ),\n")
	}
	ret.WriteString(")\n")
	return ret.String()
}

func ExecuteTemplate(c Output, temp string) {
	f, err := os.Create("code.py")
	if err != nil {
		color256.PrintHiRed("ERROR: template parsing fail")
		os.Exit(1)
	}
	t, err := template.ParseFiles(temp)
	if err != nil {
		color256.PrintHiRed("ERROR: template parsing fail")
		os.Exit(1)
	}
	t.Execute(f, c)
}
