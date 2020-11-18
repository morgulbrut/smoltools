package usbID

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func GetID(vid string, pid string, idl string) []string {
	ret := []string{"", "", "", ""}
	for _, vend := range getVendors(idl) {
		v := strings.Split(vend, "  ")[0]
		if v == vid {
			ven := strings.Split(vend, "  ")[1]
			ven = strings.Split(ven, "\n")[0]
			ret[0] = "VID:\t" + v
			ret[1] = "Vendor:\t" + ven
			prods := strings.Split(vend, "\n")[1:]
			for _, prod := range prods {
				p := strings.Split(prod, "  ")[0]
				p = strings.TrimSpace(p)
				if p == pid {
					name := strings.Split(prod, "  ")[1]
					ret[2] = "PID:\t" + p
					ret[3] = "Name:\t" + name
				}
			}
		}

	}
	return ret
}

func DownloadList(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return string(content), nil
}

func getVendors(str string) []string {
	var re = regexp.MustCompile(`(?m)(?:^[[:xdigit:]]{4}.*\n)(?:^\t[[:xdigit:]]{4}.*\n)*`)
	match := re.FindAllString(str, -1)
	return match
}
