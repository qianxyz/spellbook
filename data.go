package main

import (
	"encoding/json"
	"log"
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

var Spells []Spell

func init() {
	bytes, err := os.ReadFile("./5e-SRD-Spells.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, &Spells)
	if err != nil {
		log.Fatal(err)
	}
}
