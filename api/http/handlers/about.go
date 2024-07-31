package handlers

import (
	"net/http"
)

func AboutRenderHandler(w http.ResponseWriter, r *http.Request) {

	t, err := GetTemplate(TemplateFS, "about.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)

}
