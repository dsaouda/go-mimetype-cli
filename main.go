// Programa detecta o mimetype de um arquivo usando uma lib escrita em go que faz a detecção sem usar libs externas.
// A detecção é feita pelo que é chamado de magic numbers.
// Esses números mágicos são definidos nos primeiros bytes do arquivo.
// No link https://www.garykessler.net/library/file_sigs.html temos uma documentação como vários desses bytes mágicos
//
// Exemplo para detecção de PDF
// no link acima faça a busca: 25 50 44 46
//
// no código a implementação ficou assim:
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

type Result struct {
	Mime 		string 	`json:"mime"`
	Extension 	string 	`json:"extension"`
	Error 		string 	`json:"error"`
}

func printJson(result Result) {
	j, err := json.Marshal(result)

	if err != nil {
		panic("Problemas com json")
	}

	fmt.Print(string(j))
}

func main() {
	var file string
	flag.StringVar(&file,"file", "", "informar o nome do arquivo")
	flag.Parse()

	if strings.TrimSpace(file) == "" {
		flag.Usage()
		return
	}

	mime, extension, err := mimetype.DetectFile(file)

	if err != nil {
		r := Result{Error: err.Error()}
		printJson(r)
		return
	}

	r := Result{Mime: mime, Extension: extension}
	printJson(r)
}
