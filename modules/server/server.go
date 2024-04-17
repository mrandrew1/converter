package server

import (
	"awesomeProject1/modules/helpers"
	"awesomeProject1/modules/server/handlers"
	"log"
	"net/http"
)

func Server() {
	// Start the server
	log.Println("Starting server on port 8080")

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/generation-page/", handlers.GenerateHandler)
	http.HandleFunc("/upload-data-file/", handlers.UploadHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	helpers.TemplatesWriter(w, "main")
	//fmt.Fprintf(w, "MAIN PAGE")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	helpers.TemplatesWriter(w, "about")
	//fmt.Fprintf(w, "ABOUT PAGE")
}
