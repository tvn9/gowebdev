// template.ParseFiles example
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("base.html")
	if err != nil {
		log.Fatal(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}
}
