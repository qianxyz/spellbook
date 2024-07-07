package main

import (
	"context"
	"net/http"
	"slices"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Helper function to render a template.
// https://www.reddit.com/r/golang/comments/17d12wk/using_echo_with_ahtempl/
func render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(
			http.StatusInternalServerError,
			"failed to render response template",
		)
	}

	return nil
}

type Query struct {
	Q        string   `query:"q"`
	Class    string   `query:"class"`
	School   []string `query:"school"`
	LevelMin int      `query:"levelMin"`
	LevelMax int      `query:"levelMax"`
	Sort     string   `query:"sort"`
}

func (spell *Spell) Satisfies(query *Query) bool {
	// query contained in spell name, case-insensitive
	if !strings.Contains(
		strings.ToLower(spell.Name),
		strings.ToLower(query.Q),
	) {
		return false
	}

	// filter by class
	if query.Class != "" {
		isOfClass := false
		for _, c := range spell.Classes {
			if c.Name == query.Class {
				isOfClass = true
				break
			}
		}
		if !isOfClass {
			return false
		}
	}

	// filter by school
	if len(query.School) > 0 {
		isOfSchool := false
		for _, s := range query.School {
			if spell.School.Name == s {
				isOfSchool = true
				break
			}
		}
		if !isOfSchool {
			return false
		}
	}

	// filter by Level
	if spell.Level < query.LevelMin || spell.Level > query.LevelMax {
		return false
	}

	return true
}

func spellListHandler(ctx echo.Context) error {
	var query Query
	err := ctx.Bind(&query)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	var bookmarks []string
	cookie, err := ctx.Cookie("bookmarks")
	if err == nil {
		bookmarks = strings.Split(cookie.Value, "%2C") // comma
	}

	// filter spells by search query
	var filteredSpells []Spell
	for _, spell := range Spells {
		if spell.Satisfies(&query) {
			filteredSpells = append(filteredSpells, spell)
		}
	}

	// sort spells
	switch query.Sort {
	case "name": // data already sorted by name, do nothing
	case "-name":
		slices.Reverse(filteredSpells)
	case "level":
		slices.SortStableFunc(filteredSpells, func(a, b Spell) int {
			return a.Level - b.Level
		})
	case "-level":
		slices.SortStableFunc(filteredSpells, func(a, b Spell) int {
			return b.Level - a.Level
		})
	}

	return render(ctx, http.StatusOK, spellList(filteredSpells, bookmarks))
}

func spellDetailHandler(ctx echo.Context) error {
	spellId := ctx.Param("id")

	for _, s := range Spells {
		if s.Id == spellId {
			return render(ctx, http.StatusOK, spellDetail(s))
		}
	}

	return ctx.String(http.StatusNotFound, "spell not found")
}
