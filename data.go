package main

import (
	"encoding/json"
	"os"
)

type Spell struct {
	Id   string `json:"index"`
	Name string

	Level  int
	School struct {
		Name string
	}
	Ritual bool

	CastingTime  string `json:"casting_time"`
	Range        string
	AreaOfEffect struct {
		Size int
		Type string
	} `json:"area_of_effect"`
	Components    []string
	Material      string
	Concentration bool
	Duration      string

	Desc        []string
	HigherLevel []string `json:"higher_level"`

	Classes []struct {
		Name string
	}
}

// Parse JSON file into list of spells.
func load_spells(fileName string) ([]Spell, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var spells []Spell
	err = json.Unmarshal(bytes, &spells)
	if err != nil {
		return nil, err
	}

	return spells, nil
}
