package helpers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func TemplatesWriter(w http.ResponseWriter, page string) {
	tmpl, _ := template.ParseFiles("public/templates/" + page + ".html")
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template parsing error: %s", err)
	}
}

func ValidateMethod(w http.ResponseWriter, r *http.Request, methodName string) {
	if r.Method != methodName {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func RemoveFilesFromResults() {
	err := os.RemoveAll("./results")
	if err != nil {
		log.Default().Println(fmt.Sprintf("Error removing files: %v", err))
	}

	err2 := os.Mkdir("./results", os.ModePerm)
	if err2 != nil {
		return
	}
}
