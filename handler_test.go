package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func listSpells(t *testing.T, target string) string {
	t.Helper()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()
	if err := spellListHandler(e.NewContext(req, rec)); err != nil {
		t.Fatalf("handler error: %v", err)
	}
	return strings.ToLower(rec.Body.String())
}

func TestSpellListDefaultsToAllLevels(t *testing.T) {
	body := listSpells(t, "/spells?q=fireball")
	if !strings.Contains(body, "fireball") {
		t.Errorf("query without level params should match the level-3 spell Fireball, got: %.200s", body)
	}
}

func TestSpellListRespectsExplicitLevelMax(t *testing.T) {
	body := listSpells(t, "/spells?q=fireball&levelMax=2")
	if strings.Contains(body, "fireball") {
		t.Errorf("levelMax=2 should exclude the level-3 spell Fireball, got: %.200s", body)
	}
}
