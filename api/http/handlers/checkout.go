package handlers

import (
	"fmt"
	"net/http"

	"github.com/froostang/retail-therapy/api/user"
)

func (sm *ShoppingManager) CheckoutRenderHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: get template helper and use file
	// t, err := template.New("webpage").Parse(tmpl)

	t, err := getTemplate("checkout.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	// Construct the full file paths relative to the current working directory
	// TODO: still need to universally fix pathing
	// basePath, err := filepath.Abs("build/templates/base.html")
	// if err != nil {
	// 	log.Fatalf("Error getting absolute path for base template: %v", err)
	// }
	// formPath, err := filepath.Abs("build/templates/form.html")
	// if err != nil {
	// 	log.Fatalf("Error getting absolute path for form template: %v", err)
	// }

	// t, err := template.New("").ParseFiles(basePath, formPath)
	// if err != nil {
	// 	http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// }

	// t := template.Must(template.New("base").ParseFiles("add_container.html", "add_form.html"))
	err = t.Execute(w, ShoppingData{User: user.User{Name: "checkout man"}, Products: sm.cache.Get()})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// data := struct {
	// 	Title string
	// }{
	// 	Title: "Add Product",
	// }

	// // Render the base template with the embedded form template
	// err = t.ExecuteTemplate(w, "build/templates/base", data)
	// if err != nil {
	// 	http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// }
}
