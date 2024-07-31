package handlers

import (
	"embed"
	"html/template"
)

var TemplateFS embed.FS

func SetTemplates(fs embed.FS) {
	TemplateFS = fs
}

func GetTemplate(fs embed.FS, name string) (*template.Template, error) {

	// TODO: fix directory structure issues with templates?
	tmpl, err := template.ParseFS(fs, "build/templates/"+name)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
