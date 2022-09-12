package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"os"
	"text/template"
)

//go:embed *.tpl
var embeddedFS embed.FS

var readmeTemplate = func() *template.Template {
	b, err := embeddedFS.ReadFile("readme.tpl")
	if err != nil {
		panic(err)
	}
	return template.Must(template.New("").Parse(string(b)))
}()

type ReadmeTemplateData struct {
	Name     string
	Location string
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	input, err := getInputFromCommandLine()
	check(err)
	var buf bytes.Buffer
	err = readmeTemplate.Execute(&buf, input)
	check(err)
	fmt.Println(buf.String())
}

func getInputFromCommandLine() (*ReadmeTemplateData, error) {
	name := flag.String("name", "", "location of the image")
	location := flag.String("location", "", "location of the image")
	flag.Parse()
	if *name == "" || *location == "" {
		flag.Usage()
		return nil, fmt.Errorf("please provide required command-line flags")
	}
	input := ReadmeTemplateData{
		Name:     *name,
		Location: *location,
	}
	return &input, nil
}
