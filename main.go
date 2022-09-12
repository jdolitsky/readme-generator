package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/google/go-containerregistry/pkg/crane"
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
	input, err := buildTemplateInput()
	check(err)
	var buf bytes.Buffer
	err = readmeTemplate.Execute(&buf, input)
	check(err)
	fmt.Println(buf.String())
}

func buildTemplateInput() (*ReadmeTemplateData, error) {
	name := flag.String("name", "", "name of the image")
	location := flag.String("location", "", "location of the image")
	excludeTags := flag.String("exclude-tags", "", "comma-separated list of tags to exlcude (optional)")
	flag.Parse()
	if *name == "" || *location == "" {
		flag.Usage()
		return nil, fmt.Errorf("please provide required command-line flags")
	}

	tags, err := getTags(*location, strings.Split(*excludeTags, ","))
	if err != nil {
		return nil, err
	}

	input := ReadmeTemplateData{
		Name:     *name,
		Location: *location,
		Tags:     tags,
	}
	return &input, nil
}

func getTags(location string, excluded []string) ([]ReadmeTemplateDataTag, error) {
	allTags, err := crane.ListTags(location)
	if err != nil {
		return nil, err
	}

	// Filter out:
	// 1. cosign-related tags
	// 2. tags produced by apko-snapshot action
	// 3. tags explicitly to be excluded
	filteredTags := []string{}
	for _, tag := range allTags {
		if strings.HasPrefix(tag, "sha256-") || strings.HasPrefix(tag, "latest-") || stringInSlice(tag, excluded) {
			continue
		}
		filteredTags = append(filteredTags, tag)
	}

	// Check the digest for each tag, and collapse them into the digest they point to
	digests := map[string][]string{}
	for _, tag := range filteredTags {
		r := fmt.Sprintf("%s:%s", location, tag)
		key, err := crane.Digest(r)
		if err != nil {
			return nil, fmt.Errorf("unable to determine digest for %s: %w", r, err)
		}
		if val, ok := digests[key]; !ok {
			digests[key] = []string{tag}
		} else {
			digests[key] = append(val, tag)
		}
	}

	// Finally, build out the template data structure
	readmeTemplateDataTags := []ReadmeTemplateDataTag{}
	for digest, aliases := range digests {
		sort.Strings(aliases)
		readmeTemplateDataTags = append(readmeTemplateDataTags, ReadmeTemplateDataTag{
			Aliases:  aliases,
			Digest:   digest,
			RekorURL: fmt.Sprintf("%s/?hash=%s", RekorSearchRootURL, digest),
		})
	}

	return readmeTemplateDataTags, nil
}

// https://stackoverflow.com/a/15323988
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
