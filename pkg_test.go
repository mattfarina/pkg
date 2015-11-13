package pkg

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestPkgPkg(t *testing.T) {
	b, err := ioutil.ReadFile("./pkg.json")
	if err != nil {
		t.Fatal("Failed to read pkg.json")
	}

	p := Pkg{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		t.Fatal("Failed to unmarshal pkg.json")
	}

	if p.Name != "github.com/Masterminds/pkg" {
		t.Error("Pkg has incorrect name from pkg.json")
	}
	if p.Description != "A package definition for Go." {
		t.Error("Pkg has incorrect description from pkg.json")
	}
	if p.Keywords[0] != "vendor" {
		t.Error("Pkg has incorrect first keyword from pkg.json")
	}
	if p.Keywords[1] != "package" {
		t.Error("Pkg has incorrect second keyword from pkg.json")
	}
}

func TestPkg(t *testing.T) {
	b, err := ioutil.ReadFile("./testdata/pkg.json")
	if err != nil {
		t.Fatal("Failed to read testdata/pkg.json")
	}

	p := Pkg{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		t.Fatal("Failed to unmarshal testdata/pkg.json")
	}

	if p.Name != "github.com/mattfarina/pkg" {
		t.Error("Pkg has incorrect name from testdata/pkg.json")
	}
	if p.Image != "foo.png" {
		t.Error("Pkg has incorrect image from testdata/pkg.json")
	}
	if p.Description != "A package definition for Go." {
		t.Error("Pkg has incorrect description from testdata/pkg.json")
	}
	if p.Keywords[0] != "vendor" {
		t.Error("Pkg has incorrect first keyword from testdata/pkg.json")
	}
	if p.Keywords[1] != "package" {
		t.Error("Pkg has incorrect second keyword from testdata/pkg.json")
	}
	if p.Home != "https://example.com" {
		t.Error("Pkg has incorrect homepage from testdata/pkg.json")
	}
	if p.License != "MIT" {
		t.Error("Pkg has incorrect license from testdata/pkg.json")
	}
	if p.Owners[0].Name != "Matt Butcher" || p.Owners[0].Email != "something@example.com" || p.Owners[0].Home != "http://technosophos.com" {
		t.Error("Pkg has incorrect first owner from testdata/pkg.json")
	}
	if p.Owners[1].Name != "Matt Farina" || p.Owners[1].Email != "other@example.com" || p.Owners[1].Home != "http://mattfarina.com" {
		t.Error("Pkg has incorrect second owner from testdata/pkg.json")
	}

	if p.DevImports[0].Name != "github.com/golang/lint" || p.DevImports[0].Version != "" {
		t.Error("Pkg has incorrect first devImport from testdata/pkg.json")
	}

	if p.Imports[2].Name != "github.com/Masterminds/cookoo" || p.Imports[2].Version != "1.2.0" || p.Imports[2].Repo != "git@github.com:Masterminds/cookoo.git" || p.Imports[2].Type != "git" {
		t.Error("Pkg has incorrect third import from testdata/pkg.json")
	}

	b2, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		t.Fatal("Failed to marshal testdata/pkg.json")
	}
	if reflect.DeepEqual(b, b2) {
		t.Error("Pkg testdata/pkg.json marshaled not same as original")
	}
}
