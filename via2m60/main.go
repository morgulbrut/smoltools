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
	config := converter.ProcessJSON("dz60_via.json")
	converter.PrintConfig(config)
}

// func main() {
// 	logo()
// 	var template string
// 	var pf string
// 	var payloads []string
// 	flag.StringVar(&template, "t", "code_template.py", "Template for code.py generation, see examples in the repository")
// 	flag.StringVar(&pf, "p", "", "DuckyScript DuckyScript payloads, get included into yout template as payload_0() to payload_n().")
// 	flag.Parse()
// 	color256.PrintHiGreen(template)
// 	color256.PrintHiGreen(pf)

// 	counter := 0
// 	pfls := strings.Split(pf, " ")
// 	for counter < len(pfls) {
// 		payloads = append(payloads, readFile(pfls[counter], counter))
// 		counter++
// 	}

// 	executeTemplate(strings.Join(payloads, "\n\n\n"), template)
// }

// func readFile(f string, c int) string {

// 	var payload []string
// 	payload = append(payload, fmt.Sprintf("def payload_%d():", c))

// 	if len(os.Args) <= 1 {
// 		color256.PrintHiRed("ERROR: Exepted JSON file")
// 		os.Exit(1)
// 	}

// 	color256.PrintHiGreen("Compiliing JSON")

// 	file, err := os.Open(f)
// 	if err != nil {
// 		color256.PrintHiRed("ERROR: File not found")
// 		os.Exit(1)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		payload = append(payload, processLine(scanner.Text()))
// 	}

// 	if err := scanner.Err(); err != nil {
// 		color256.PrintHiRed("ERROR: Could not read file")
// 		os.Exit(1)
// 	}
// 	return strings.Join(payload, "\n")
// }

// func processLine(s string) string {
// 	params := strings.Split(s, " ")
// 	cmd, params := params[0], params[1:]
// 	cmd = strings.ToUpper(cmd)
// 	switch cmd {
// 	case "REM":
// 		return "    # " + strings.Join(params, " ")
// 	case "STRING":
// 		return "    keyboard_layout.write(\"" + strings.Join(params, " ") + "\")"
// 	case "DELAY":
// 		t, err := strconv.ParseFloat(params[0], 64)
// 		if err != nil {
// 			color256.PrintHiRed("ERROR: DELAY: Missing param")
// 			os.Exit(1)
// 		}
// 		return fmt.Sprintf("    time.sleep(%0.1f)", t/1000)
// 	case "GUI", "SHIFT", "ALT", "CTRL":
// 		key := strings.ToUpper(params[0])
// 		switch key {
// 		case "SHIFT", "ALT", "CTRL":
// 			key2 := strings.ToUpper(params[1])
// 			return "    keyboard.press(Keycode." + cmd + ", Keycode." + key + ", Keycode." + key2 + ")\n    keyboard.release_all()"
// 		default:
// 			return "    keyboard.press(Keycode." + cmd + ", Keycode." + key + ")\n    keyboard.release_all()"
// 		}
// 	case "DELETE", "HOME", "INSERT", "PAGE_UP", "PAGE_DOWN", "UP_ARROW", "DOWN_ARROW", "LEFT_ARROW", "RIGHT_ARROW", "TAB",
// 		"END", "ESCAPE", "ENTER":
// 		return "    keyboard.press(Keycode." + cmd + ")\n    keyboard.release_all()"
// 	default:
// 		return "    # UNDEFINED: " + s
// 	}
// }

// func executeTemplate(s, temp string) {
// 	color256.PrintHiYellow(s)
// 	f, err := os.Create("code_via.py")
// 	if err != nil {
// 		color256.PrintHiRed("ERROR: template parsing fail")
// 		os.Exit(1)
// 	}
// 	t, err := template.ParseFiles(temp)
// 	if err != nil {
// 		color256.PrintHiRed("ERROR: template parsing fail")
// 		os.Exit(1)
// 	}
// 	t.Execute(f, s)
// }
