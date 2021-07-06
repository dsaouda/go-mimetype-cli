// Program detects the mimetype of a file using a lib written in go which does the detection without using external libs.
// Detection is done by what are called magic numbers.
// These magic numbers are defined in the first bytes of the file.
// At the link https://www.garykessler.net/library/file_sigs.html we have documentation on several of these magic bytes
//
// Example for PDF detection
// on the link above do the search: 25 50 44 46
//
// in the code the implementation looked like this:
//
// func Pdf(in []byte) bool {
// 		return len(in) > 4 && bytes.Equal(in[:4], []byte{0x25, 0x50, 0x44, 0x46})
// }

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"strings"
)

type ok struct {
	Mime 		string 	`json:"mime"`
	Extension 	string 	`json:"extension"`
}

type fail struct {
	Error 		string 	`json:"error"`
}

func printJson(result interface{}) {
	j, err := json.Marshal(result)

	if err != nil {
		panic("json error")
	}

	fmt.Print(string(j))
}

func main() {
	var file string
	flag.StringVar(&file,"file", "", "type filename")
	flag.Parse()

	if strings.TrimSpace(file) == "" {
		flag.Usage()
		return
	}

	mime, err := mimetype.DetectFile(file)

	if err != nil {
		r := fail{Error: err.Error()}
		printJson(r)
		return
	}

	r := ok{Mime:mime.String(), Extension: mime.Extension()}
	printJson(r)
}
