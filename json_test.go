package mirror

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"testing"

	srv "github.com/google/zoekt/cmd/zoekt-indexserver"
)

func TestJSON(t *testing.T) {
	j, err := ioutil.ReadFile("mirror.json")
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}

	result := []srv.ConfigEntry{}
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
