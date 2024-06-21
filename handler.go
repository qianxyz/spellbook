package main

import (
	"context"
	"net/http"

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
func index(ctx echo.Context) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return render(ctx, http.StatusOK, page(Spells))
}
