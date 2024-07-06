package main

import (
	"html/template"
	"os"
)

func main() {
	// fmt.Println("hell go")

	// fmt.Printf("Hello %s\n", "World")
	// fmt.Printf("Hello %d\n", 343)

	type InputData struct {
		Title     string
		Body      string
		MenuItems []string
	}

	file, err := os.ReadFile("./template/home.html")
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("sample").Parse(string(file)))

	outfile, err := os.Create("./output.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outfile, InputData{
		Title:     "Sathvik Racha | Resume Site",
		Body:      "This my resume",
		MenuItems: []string{"fist", "second", "third"},
	})
	if err != nil {
		panic(err)
	}
}
