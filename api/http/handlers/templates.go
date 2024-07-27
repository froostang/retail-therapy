package handlers

import (
	"html/template"
	"os"
	"path/filepath"
)

func getTemplate(name string) (*template.Template, error) {
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exPath := filepath.Dir(ex)

	// TODO: fix directory structure issues with templates
	tmpl, err := template.ParseFiles(exPath + "/templates/" + name)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
