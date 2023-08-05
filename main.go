// main.go

package main

import (
	"net/http"
)

func main() {
	items := []Item{
		{1, "Item A", 10.0},
		{2, "Item B", 15.0},
		{3, "Item C", 20.0},
	}

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan data dari model (items)
	data := items

	// Merender tampilan menggunakan view (template)
	renderTemplate(w, "index", data)
}
