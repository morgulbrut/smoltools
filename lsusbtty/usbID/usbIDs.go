package usbID

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func GetID(vid string, pid string) string {

	for _, vend := range getVendors(downloadList("http://www.linux-usb.org/usb.ids")) {
		v := strings.Split(vend, "  ")[0]
		ven := strings.Split(vend, "  ")[1]
		if v == vid {
			fmt.Printf("\t\tVID: %s, MAN: %s\n", v, ven)
		}
	}

	return ""
}

func downloadList(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return string(content)
}

func getVendors(str string) []string {
	var re = regexp.MustCompile(`(?m)(?:^[[:xdigit:]]{4}.*\n)(?:^\t[[:xdigit:]]{4}.*\n)*`)
	match := re.FindAllString(str, -1)
	return match
}
