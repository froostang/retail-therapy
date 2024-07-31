package handlers

import (
	"net/http"
)

func AdderRenderHandler(w http.ResponseWriter, r *http.Request) {

	t, err := getTemplate("add_updated.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)

}
