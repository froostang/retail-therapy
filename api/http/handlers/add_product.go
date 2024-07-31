package handlers

import (
	"net/http"
)

func AdderRenderHandler(w http.ResponseWriter, r *http.Request) {

	t, err := GetTemplate(TemplateFS, "add_product.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)

}
