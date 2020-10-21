package main

import (
	"github.com/nobonobo/spago"
	"github.com/nobonobo/spago-spectre/forms"
)

// Top ...
type Top struct {
	spago.Core
}

// Render ...
func (c *Top) Render() spago.HTML {
	return spago.Tag("body",
		spago.Tag("main",
			spago.Tag("form",
				spago.ClassMap{
					"card":       true,
					"p-centered": true,
					"col-6":      true,
					"col-lg-8":   true,
					"col-md-9":   true,
				},
				spago.Tag("div", spago.ClassMap{"card-body": true},
					spago.C(
						&forms.InputForm{Config: forms.Config{Name: "name", Label: "名前"}, Type: "text"},
						&forms.InputForm{Config: forms.Config{Name: "password", Label: "パスワード"}, Type: "text"},
						&forms.InputForm{Config: forms.Config{Name: "datetime", Label: "日付"}, Type: "datetime-local"},
						&forms.InputForm{Config: forms.Config{Name: "email", Label: "メールアドレス"}, Type: "email"},
						&forms.InputForm{Config: forms.Config{Name: "url", Label: "ホームページ"}, Type: "url"},
					),
					spago.ClassMap{"form-group": true},
					spago.Tag("label", spago.T("情報")),
					spago.C(
						&forms.CheckboxForm{Config: forms.Config{Name: "registered", Label: "登録"}, Type: "switch"},
					),
					spago.Tag("div",
						spago.ClassMap{"form-group": true},
						spago.Tag("label", spago.T("性別")),
						spago.Tag("div",
							spago.ClassMap{"form-group": true},
							spago.C(
								&forms.CheckboxForm{Config: forms.Config{Name: "gender", Label: "申告しない", Inline: true}, Type: "radio", Value: true},
								&forms.CheckboxForm{Config: forms.Config{Name: "gender", Label: "男性", Inline: true}, Type: "radio"},
								&forms.CheckboxForm{Config: forms.Config{Name: "gender", Label: "女性", Inline: true}, Type: "radio"},
							),
						),
					),
				),
				spago.Tag("div",
					spago.ClassMap{"card-footer": true},
					spago.Tag("button",
						spago.AttrMap{
							"type":  "submit",
							"class": "btn",
						},
						spago.T("Submit"),
					),
				),
			),
		),
	)
}

func main() {
	spago.AddStylesheet("https://unpkg.com/spectre.css/dist/spectre.min.css")
	spago.RenderBody(&Top{})
	select {}
}
