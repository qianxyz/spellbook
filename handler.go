package main

import (
	"context"
	"net/http"
	"strconv"
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

func spellListHandler(ctx echo.Context) error {
	query := ctx.QueryParam("q")
	class := ctx.QueryParam("class")
	schools := ctx.QueryParams()["school"]

	level := ctx.QueryParam("level")
	levels := strings.Split(level, ",")
	levelMin, err := strconv.Atoi(levels[0])
	if err != nil {
		levelMin = 0
	}
	levelMax, err := strconv.Atoi(levels[1])
	if err != nil {
		levelMax = 9
	}

	// filter spells by search query
	var filteredSpells []Spell
	for _, spell := range Spells {
		// query contained in spell name, case-insensitive
		if !strings.Contains(
			strings.ToLower(spell.Name),
			strings.ToLower(query),
		) {
			continue
		}

		// filter by class
		if class != "" {
			hasClass := false
			for _, c := range spell.Classes {
				if c.Name == class {
					hasClass = true
					break
				}
			}
			if !hasClass {
				continue
			}
		}

		// filter by school
		if len(schools) > 0 {
			hasSchool := false
			for _, s := range schools {
				if spell.School.Name == s {
					hasSchool = true
					break
				}
			}
			if !hasSchool {
				continue
			}
		}

		// filter by level
		if spell.Level < levelMin || spell.Level > levelMax {
			continue
		}

		filteredSpells = append(filteredSpells, spell)
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
