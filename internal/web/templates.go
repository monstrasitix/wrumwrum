package web

import (
	"log"
	"text/template"
)

const (
	layoutBase = "./templates/layout/base.go.tpl"
)

var HTML = map[string]*template.Template{}

func Setup() {
	temp := template.New("base").Funcs(template.FuncMap{
		"tableAlign": func(value string) string {
			if value == "numeric" {
				return "-numeric"
			}

			return ""
		},
		"slice": func(size int) []int {
			return make([]int, size)
		},
	})

	temp, err := temp.ParseGlob("./templates/partial/**.go.tpl")
	if err != nil {
		log.Fatal(err)
	}

	HTML["home"] = page(temp, layoutBase, "./templates/page/home.go.tpl")
	HTML["table"] = page(temp, layoutBase, "./templates/page/table.go.tpl")
	HTML["charts"] = page(temp, layoutBase, "./templates/page/charts.go.tpl")
}

func page(temp *template.Template, paths ...string) *template.Template {
	page, err := template.Must(temp.Clone()).ParseFiles(paths...)
	if err != nil {
		log.Fatal(err)
	}

	return page
}
