package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//Parsing the file
	tpl, err := template.ParseFiles("christmas-card.gohtml")
	if err != nil {
		log.Panic(err)
	}

	friends := []string{"Rivan", "Jhouma", "Asus", "Onit", "Wahyu"}

	//Executing the file
	//ExecuteTemplate give the output of type error
	err = tpl.ExecuteTemplate(os.Stdout, "christmas-card.gohtml", friends)
	if err != nil {
		log.Printf("%s\r\n", err)
	}

}
