package internal

import "github.com/maxence-charriere/go-app/v7/pkg/app"

func NewWebHandler() *app.Handler {
	return &app.Handler{
		Name:               "minim",
		Title:              "minim",
		Description:        "a note app",
		Icon:               app.Icon{
			Default: "/web/favicon.svg",
		},

		Scripts:            []string{
			"/web/lib/less@3.12.2/less.min.js",
		},
		Styles:             []string{
			"/web/lib/purecss@2.0.3/pure-min.css",
			"/web/lib/purecss@2.0.3/grids-responsive-min.css",
		},
		RawHeaders: []string{
			`<link rel="stylesheet/less" type="text/css" href="/web/styles.less" />`,
		},
	}
}
