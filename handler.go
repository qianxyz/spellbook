package main

import (
	"context"
	"net/http"
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
	Q      string
	Class  string
	School []string
	Level  []int
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
	if spell.Level < query.Level[0] || spell.Level > query.Level[1] {
		return false
	}

	return true
}

func spellListHandler(ctx echo.Context) error {
	var query Query
	err := echo.QueryParamsBinder(ctx).
		String("q", &query.Q).
		String("class", &query.Class).
		Strings("school", &query.School).
		BindWithDelimiter("level", &query.Level, ",").
		BindError()
	if err != nil || len(query.Level) < 2 {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	// filter spells by search query
	var filteredSpells []Spell
	for _, spell := range Spells {
		if spell.Satisfies(&query) {
			filteredSpells = append(filteredSpells, spell)
		}
	}

	return render(ctx, http.StatusOK, spellList(filteredSpells))
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
