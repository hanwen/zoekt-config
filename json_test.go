package mirror

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"testing"
)

type configEntry struct {
	GithubUser string
	GitilesURL string
	CGitURL    string
	Name       string
	Exclude    string
}

func TestJSON(t *testing.T) {
	j, err := ioutil.ReadFile("mirror.json")
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}

	result := []configEntry{}
	if err := json.Unmarshal(j, &result); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	for _, r := range result {
		if _, err := regexp.Compile(r.Name); err != nil {
			t.Fatalf("compile(%q): %v", r.Name, err)
		}
		if _, err := regexp.Compile(r.Exclude); err != nil {
			t.Fatalf("compile(%q): %v", r.Name, err)
		}
	}
}
