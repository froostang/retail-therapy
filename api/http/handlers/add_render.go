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
    <script>
        document.addEventListener('DOMContentLoaded', function () {
            document.getElementById('addProductForm').addEventListener('submit', function (event) {
                event.preventDefault(); // Prevent the default form submission

                // Collect the input data
                const inputField = document.querySelector('input[name="location"]');
                const data = {
                    url: inputField.value
                };

                // Send a POST request with JSON body using fetch
                fetch('/add', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(data)
                })
                .then(response => response.text())
                .then(text => {
                    document.getElementById('response').innerHTML = text;
                })
                .catch(error => {
                    console.error('Error:', error);
                });
            });
        });
    </script>
</head>
<body>
    <h1>Add Product Location</h1>
    <form id="addProductForm">
        <input type="text" name="location" placeholder="Enter location" required />
        <button type="submit">Add</button>
    </form>
    <div id="response"></div>
</body>
</html>
`
	// TODO: get template helper and use file
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
