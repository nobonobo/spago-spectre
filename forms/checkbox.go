package forms

import (
	"github.com/nobonobo/spago"
)

// CheckboxForm ...
type CheckboxForm struct {
	spago.Core
	Config Config
	Type   string
	Value  bool
}

// Render ...
func (c *CheckboxForm) Render() spago.HTML {
	var typ, class string
	switch c.Type {
	default:
		typ = "checkbox"
		class = "form-checkbox"
	case "switch":
		typ = "checkbox"
		class = "form-switch"
	case "radio":
		typ = "radio"
		class = "form-radio"
	}
	return spago.Tag("label",
		spago.ClassMap{
			class:         true,
			"form-inline": c.Config.Inline,
			"is-success":  c.Config.IsSuccess,
			"is-error":    c.Config.IsError,
		},
		spago.Tag("input",
			append([]spago.Markup{
				spago.ClassMap{
					"disabled": c.Config.Disabled,
				},
				spago.A("type", typ),
				spago.A("id", c.Config.ID),
				spago.A("name", c.Config.Name),
				spago.If(c.Value, spago.A("checked", true)),
			}, c.Config.Markups...)...,
		),
		spago.Tag("i", spago.ClassMap{"form-icon": true}),
		spago.T(c.Config.Label),
	)
}
