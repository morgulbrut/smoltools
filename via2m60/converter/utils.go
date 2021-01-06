package converter

import (
	"io/ioutil"
	"os"

	"github.com/morgulbrut/color256"
)

func readfile(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		color256.PrintHiRed(err.Error())
	}
	defer file.Close()
	val, err := ioutil.ReadAll(file)
	if err != nil {
		color256.PrintHiRed(err.Error())
	}
	return val
}
