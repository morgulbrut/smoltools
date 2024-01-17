package main

import (
	"fmt"
	"os"
	"time"

	"github.com/acarl005/stripansi"
	"github.com/morgulbrut/color256"
	"gopkg.in/mcuadros/go-syslog.v2"
	"gopkg.in/mcuadros/go-syslog.v2/format"
)

var LOGLEVEL = map[int]string{
	0: color256.BgHiRed("EMERGENCY:"),
	1: color256.BgRed("ALERT   :"),
	2: color256.BgYellow("CRITICAL:"),
	3: color256.Red(color256.Bold("ERROR   :")),
	4: color256.Orange(color256.Bold("WARNING :")),
	5: color256.Yellow(color256.Bold("NOTICE  :")),
	6: color256.Magenta(color256.Bold("INFO    :")),
	7: color256.Green("DEBUG   :"),
	8: color256.Cyan("TRACE   :"),
}

func logo() {
	logo := `                                                                        
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓ 
┃                                     __ _______ _______                   ┃░
┃                  .-----.--.--.-----|  |   _   |   _   |                  ┃░
┃                  |__ --|  |  |__ --|  |.  |___|.  |   |                  ┃░
┃                  |_____|___  |_____|__|.  |   |.  |   |                  ┃░
┃                        |_____|        |:  1   |:  1   |                  ┃░
┃                                       |::.. . |::.. . |                  ┃░
┃                                       '-------'-------'                  ┃░
┃                  Simple syslog server in go                              ┃░
┃                  v 0.1                                                   ┃░
┃                                                                          ┃░
┃                                                                          ┃░
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛░
 ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░`

	fmt.Println(color256.Magenta(logo))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func output(logParts format.LogParts) string {
	t := logParts["timestamp"].(time.Time)
	ts := fmt.Sprintf(t.Format(time.RFC3339))
	hn := color256.BgHiBlue(logParts["client"].(string))
	return fmt.Sprintf("%s %s %s %s", color256.Faint(ts), hn, LOGLEVEL[logParts["severity"].(int)], logParts["content"])
}

func main() {
	logo()
	fn := "syslgo.log"
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_APPEND, 0644)
	check(err)
	defer f.Close()

	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	server.ListenUDP("0.0.0.0:514")
	server.ListenTCP("0.0.0.0:514")
	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			o := output(logParts)
			fmt.Println(o)
			f.WriteString(stripansi.Strip(o) + "\n")
		}
	}(channel)

	server.Wait()
}
