package main

import (
	"context"
	"log"

	"github.com/shurcooL/graphql"
)

type Spell struct {
	Id   string `graphql:"index"`
	Name string

	Level  int
	School struct {
		Name string
	}
	Ritual bool

	CastingTime  string `graphql:"casting_time"`
	Range        string
	AreaOfEffect struct {
		Size int
		Type string
	} `graphql:"area_of_effect"`
	Components    []string
	Material      string
	Concentration bool
	Duration      string

	Desc        []string
	HigherLevel []string `graphql:"higher_level"`

	Classes []struct {
		Name string
	}
}

var Spells []Spell

func init() {
	q := struct {
		Spells []Spell `graphql:"spells(limit: 500)"`
	}{}

	url := "http://www.dnd5eapi.co/graphql"
	client := graphql.NewClient(url, nil)

	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		log.Fatal(err)
	}

	Spells = q.Spells
}
