package handlers

import (
	"html/template"
	"net/http"
)

func AdderRenderHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Add Product</title>
        <script src="https://unpkg.com/htmx.org@1.9.4/dist/htmx.min.js"></script>
    </head>
    <body>
        <h1>Add Product Location</h1>
        <form id="addProductForm" hx-post="/add" hx-target="#response" hx-swap="innerHTML">
            <input type="text" name="location" placeholder="Enter location" required />
            <button type="submit">Add</button>
        </form>
        <div id="response"></div>
    </body>
    </html>`
	// TODO: get template helper and use file
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
