package forms

import "github.com/nobonobo/spago"

// Config ...
type Config struct {
	Label     string
	ID        string
	Name      string
	Disabled  bool
	IsSuccess bool
	IsError   bool
	Hint      string
	Inline    bool
	Markups   []spago.Markup
}
