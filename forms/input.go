package forms

import (
	"github.com/nobonobo/spago"
)

// InputForm ...
type InputForm struct {
	spago.Core
	Config Config
	Type   string
	Value  string
}

// Render ...
func (c *InputForm) Render() spago.HTML {
	return spago.Tag("div",
		spago.ClassMap{
			"form-group":  true,
			"has-success": c.Config.IsSuccess,
			"has-error":   c.Config.IsError,
		},
		spago.Tag("label",
			spago.ClassMap{"form-label": true},
			spago.If(len(c.Config.ID) > 0, spago.A("for", c.Config.ID)),
			spago.T(c.Config.Label),
		),
		spago.Tag("input",
			append([]spago.Markup{
				spago.ClassMap{
					"form-input": true,
					"disabled":   c.Config.Disabled,
				},
				spago.A("type", c.Type),
				spago.If(len(c.Config.ID) > 0, spago.A("id", c.Config.ID)),
				spago.A("name", c.Config.Name),
				spago.If(len(c.Value) > 0, spago.T(c.Value)),
			}, c.Config.Markups...)...,
		),
		spago.If(len(c.Config.Hint) > 0, spago.Tag("p",
			spago.ClassMap{
				"form-input-hint": true,
			},
			spago.T(c.Config.Hint),
		)),
	)
}
