package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"

	"github.com/google/go-containerregistry/pkg/crane"
	ispec "github.com/opencontainers/image-spec/specs-go/v1"
)

const RekorSearchRootURL = "https://rekor.tlog.dev"

//go:embed *.tpl
var embeddedFS embed.FS

var templateFuncs = template.FuncMap{
	"host": func(location string) string {
		return strings.Split(location, "/")[0]
	},
	"formatList": func(items []string) string {
		tmp := []string{}
		for _, item := range items {
			tmp = append(tmp, fmt.Sprintf("`%s`", item))
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
	Repo          string
	Name          string
	Description   string
	Location      string
	UsageMarkdown string
	CosignOutput  string
	UsesMelange   bool
	Tags          []ReadmeTemplateDataTag
}

type ReadmeTemplateDataTag struct {
	Aliases  []string
	Archs    []string
	Digest   string
	RekorURL string
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
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
	repo := flag.String("repo", "", "GitHub repo URL")
	name := flag.String("name", "", "Name of the image")
	location := flag.String("location", "", "Location of the image")
	description := flag.String("description", "", "Description of the image")
	excludeTags := flag.String("exclude-tags", "", "Comma-separated list of tags to exlcude (optional)")
	flag.Parse()
	if *repo == "" || *name == "" || *location == "" || *description == "" {
		flag.Usage()
		return nil, fmt.Errorf("please provide required command-line flags")
	}

	cosignOutput, err := cosignVerifyImage(fmt.Sprintf("%s:latest", *location))
	if err != nil {
		return nil, err
	}

	tags, err := getTags(*location, strings.Split(*excludeTags, ","))
	if err != nil {
		return nil, err
	}

	rawUsageURL := fmt.Sprintf("%s/main/USAGE.md", strings.Replace(*repo, "https://github.com", "https://raw.githubusercontent.com", 1))
	usageMarkdown := fetchRemoteFileContents(rawUsageURL)
	usesMelange := remoteFileExists(fmt.Sprintf("%s/blob/main/melange.yaml", *repo))

	input := ReadmeTemplateData{
		Repo:          *repo,
		Name:          *name,
		Description:   *description,
		Location:      *location,
		UsageMarkdown: usageMarkdown,
		UsesMelange:   usesMelange,
		CosignOutput:  cosignOutput,
		Tags:          tags,
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

		// determine the supported architectures
		digestLocation := fmt.Sprintf("%s@%s", location, digest)
		b, err := crane.Manifest(digestLocation)
		if err != nil {
			return nil, fmt.Errorf("unable to fetch manifest at %s: %w", digestLocation, err)
		}
		var index ispec.Index
		if err := json.Unmarshal(b, &index); err != nil {
			return nil, fmt.Errorf("unable to convert manifest to OCI Index at %s: %w", digestLocation, err)
		}
		archs := []string{}
		for _, m := range index.Manifests {
			if m.Platform != nil {
				archs = append(archs, m.Platform.Architecture+m.Platform.Variant)
			}
		}
		sort.Strings(archs)

		readmeTemplateDataTags = append(readmeTemplateDataTags, ReadmeTemplateDataTag{
			Aliases:  aliases,
			Archs:    archs,
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

// https://stackoverflow.com/a/42691977
func remoteFileExists(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		return false
	}
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func fetchRemoteFileContents(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(b)
}

func cosignVerifyImage(ref string) (string, error) {
	cmd := fmt.Sprintf("COSIGN_EXPERIMENTAL=1 cosign verify %s | jq", ref)
	b, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("running \"%s\" failed: %w", cmd, err)
	}
	return strings.TrimSpace(string(b)), nil
}
