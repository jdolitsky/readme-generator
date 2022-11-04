package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

type VersionsManifest struct {
	Alternatives []VersionsManifestAlternative `yaml:"alternatives"`
	Custom       bool                          `yaml:"custom"`
	Versions     []VersionsManifestVersion     `yaml:"versions"`
}

type VersionsManifestAlternative struct {
	Name string `yaml:"name"`
}

type VersionsManifestVersion struct {
	Version string                        `yaml:"version"`
	Tags    []string                      `yaml:"tags"`
	Source  VersionsManifestVersionSource `yaml:"source"`
}

type VersionsManifestVersionSource struct {
	Uri    string `yaml:"uri"`
	Sha256 string `yaml:"sha256"`
}

// go build -o repo-enforcer . && (cd testdata/myimage-custom/ && ../../repo-enforcer)
func main() {
	f, err := os.Open("versions.yaml")
	check(err)
	defer f.Close()

	b, err := io.ReadAll(f)
	check(err)

	var m VersionsManifest
	err = yaml.Unmarshal(b, &m)
	check(err)

	fmt.Println(m.Versions)
}
