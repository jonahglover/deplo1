package recipe

import (
	"encoding/json"
	"io/ioutil"
)

// Recipe definition.
type Recipe struct {
	Name     string `json:"name"`
	Filename string `json:"filename"`
}

// New Recipe loaded from file
func New(file string) (*Recipe, error) {
	r := new(Recipe)

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(b, &r)
	return r, nil
}
