package mirror

import (
	"encoding/json"
	"io/ioutil"
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
}
