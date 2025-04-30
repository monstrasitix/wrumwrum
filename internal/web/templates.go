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
	temp, err := template.ParseGlob("./templates/partial/**.go.tpl")
	if err != nil {
		log.Fatal("Could not load partials")
	}

	HTML["home"] = page(temp, layoutBase, "./templates/page/home.go.tpl")
	HTML["login"] = page(temp, layoutBase, "./templates/page/login.go.tpl")
}

func page(temp *template.Template, paths ...string) *template.Template {
	page, err := template.Must(temp.Clone()).ParseFiles(paths...)
	if err != nil {
		log.Fatal(err)
	}

	return page
}
