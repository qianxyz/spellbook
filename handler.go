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

// Handler for the index page.
func indexHandler(ctx echo.Context) error {
	return render(ctx, http.StatusOK, index())
}

func spellListHandler(ctx echo.Context) error {
	search := ctx.QueryParam("search")

	// filter spells by search query
	var filteredSpells []Spell
	for _, spell := range Spells {
		// query contained in spell name, case-insensitive
		if strings.Contains(
			strings.ToLower(spell.Name),
			strings.ToLower(search),
		) {
			filteredSpells = append(filteredSpells, spell)
		}
	}

	return render(ctx, http.StatusOK, spellList(filteredSpells))
}
