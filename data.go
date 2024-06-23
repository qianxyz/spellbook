package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
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

// Lowercase the first letter of a string.
func lowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// The level and school of the spell, and whether it is a ritual.
func (spell *Spell) LevelSchoolString() string {
	var s string

	if spell.Level == 0 {
		s = spell.School.Name + " cantrip"
	} else {
		switch spell.Level {
		case 1:
			s = "1st"
		case 2:
			s = "2nd"
		case 3:
			s = "3rd"
		default:
			s = strconv.Itoa(spell.Level) + "th"
		}
		s += "-level " + strings.ToLower(spell.School.Name)
	}
	if spell.Ritual {
		s += " (ritual)"
	}

	return s
}

// The displayed casting time of the spell.
func (spell *Spell) CastingTimeString() string {
	return spell.CastingTime
}

// The displayed range of the spell, including area of effect if applicable.
func (spell *Spell) RangeString() string {
	s := spell.Range

	if spell.Range == "Self" && spell.AreaOfEffect.Type != "" {
		s += " ("
		s += strconv.Itoa(spell.AreaOfEffect.Size) + "-foot "
		s += spell.AreaOfEffect.Type
		s += ")"
	}

	return s
}

// The displayed components of the spell, including material if applicable.
func (spell *Spell) ComponentsString() string {
	s := strings.Join(spell.Components, ", ")

	if spell.Material != "" {
		s += " (" + strings.TrimSuffix(lowerFirst(spell.Material), ".") + ")"
	}

	return s
}

// The displayed duration of the spell, including concentration if applicable.
func (spell *Spell) DurationString() string {
	if spell.Concentration {
		return "Concentration, " + lowerFirst(spell.Duration)
	} else {
		return spell.Duration
	}
}

// Used to convert the spell's description to HTML.
// See examples at https://github.com/gomarkdown/markdown
func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

// The description of the spell.
func (spell *Spell) DescriptionString() string {
	var s string
	for _, d := range spell.Desc {
		if strings.HasSuffix(s, "|") && strings.HasPrefix(d, "|") {
			// tables
			s += "\n" + d
		} else if strings.HasPrefix(d, "-") {
			// lists
			s += "\n" + d
		} else {
			s += "\n\n" + d
		}
	}

	s = string(mdToHTML([]byte(s)))

	// add bootstrap class for <table>
	// HACK: Some better options for styling:
	// - Extend Bootstrap class (requires Sass)
	// - Add class to elements with JavaScript
	// - Configure the renderer
	s = strings.ReplaceAll(s, "<table>", `<table class="table table-striped">`)

	return s
}

// The effect of casting the spell at a higher level.
// NOTE: In the current data the array has at most one element.
func (spell *Spell) HigherLevelString() string {
	if len(spell.HigherLevel) > 0 {
		return spell.HigherLevel[0]
	} else {
		return ""
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
