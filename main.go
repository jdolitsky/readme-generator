package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

const RekorSearchRootURL = "https://rekor.tlog.dev"

//go:embed *.tpl
var embeddedFS embed.FS

var templateFuncs = template.FuncMap{
	"formatTagAliases": func(tagAliases []string) string {
		tmp := []string{}
		for _, alias := range tagAliases {
			tmp = append(tmp, fmt.Sprintf("`%s`", alias))
		}
		return strings.Join(tmp, " ")
	},
}

var readmeTemplate = func() *template.Template {
	b, err := embeddedFS.ReadFile("readme.tpl")
	if err != nil {
		panic(err)
	}
	return template.Must(template.New("").Funcs(templateFuncs).Parse(string(b)))
}()

type ReadmeTemplateData struct {
	Name     string
	Location string
	Tags     []ReadmeTemplateDataTag
}

type ReadmeTemplateDataTag struct {
	Aliases  []string
	Digest   string
	RekorURL string
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
		Tags: []ReadmeTemplateDataTag{
			{
				Aliases:  []string{"a", "b", "c"},
				Digest:   "sha256:abc",
				RekorURL: fmt.Sprintf("%s/?digest=sha256:abc", RekorSearchRootURL),
			},
			{
				Aliases:  []string{"x", "y", "z"},
				Digest:   "sha256:xyz",
				RekorURL: fmt.Sprintf("%s/?digest=sha256:xyz", RekorSearchRootURL),
			},
		},
	}
	return &input, nil
}
