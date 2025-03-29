package services

import (
	"embed"
	"html/template"
	"log"
)

//go:embed "templates/*"
var template_files embed.FS

var Templates *template.Template

func init()  {
	tpl, err := template.New("").ParseFS(template_files, "templates/*.html")
	if err != nil {
		log.Fatalf("Fuck: %s", err)
	}
	Templates = tpl
}
