package handlers

import (
	"awesomeProject1/modules/createWordFiles"
	"awesomeProject1/modules/helpers"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const MaxUploadSize = 10 * 1024 * 1024

type GenerationRequest struct {
	Name string `json:"name"`
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method == "POST" {
		// Parse the JSON request body into a struct
		var req GenerationRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad GenerationRequest, invalid JSON "+err.Error(), http.StatusBadRequest)
			return
		}
		// Create the JSON response
		res := Response{
			Message: "Hello " + req.Name,
		}

		// Encode the JSON response
		jsonResponse, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to "application/json"
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the http.ResponseWriter
		w.Write(jsonResponse)
	} else if r.Method == "GET" {
		helpers.TemplatesWriter(w, "generatePage")
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
func UploadDocxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	helpers.ValidateMethod(w, r, "POST")
	//checkFolder()
	//todo get name from request
	saved := saveExampleAndDataFiles(w, r)
	if saved {
		processFiles(w, r)
	} else {
		http.Error(w, "not all files uploaded, please load another one time", http.StatusBadRequest)
		//return Response{Message: "not all files uploaded"}
	}
}
func checkFolder() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
func processFiles(w http.ResponseWriter, r *http.Request) {
	helpers.RemoveFilesFromResults()

	//generate docs and send them to user
	_, fileHeader1, _ := r.FormFile("csv_file")
	_, fileHeader2, _ := r.FormFile("docx_file")
	//_, newFileName, _ := r.FormFile("string_input")
	newFileName := strings.Trim(fileHeader2.Filename, ".docx")
	createWordFiles.GenerateWordFilesFromCsV(fileHeader1.Filename, fileHeader2.Filename, newFileName)
	files := readFolder("./results")

	helpers.CreateZip(files, newFileName)

	sendToUser(w, r, "./to_download/result_documents.zip", "result_documents.zip")
}

func saveExampleAndDataFiles(w http.ResponseWriter, r *http.Request) bool {
	savedCount := 0
	if validateFile(w, r, "text/plain; charset=utf-8", "csv_file") {
		saved := saveFile(w, r, "csv_file", "./uploads/data")
		if saved {
			savedCount++
		}
	}
	if validateFile(w, r, "application/zip", "docx_file") {
		saved := saveFile(w, r, "docx_file", "./uploads/examples")
		if saved {
			savedCount++
		}
	}

	if savedCount == 2 {
		return true
	}
	return false
}
func sendToUser(w http.ResponseWriter, r *http.Request, filePath string, filename string) {
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, filePath)
}

func readFolder(dirname string) []string {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(fmt.Sprintf("read folder error: %s", err.Error()))
	}
	infos := make([]string, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			log.Fatal(fmt.Sprintf("read folder error 2: %s", err.Error()))
		}
		infos = append(infos, info.Name())
	}
	return infos
}

func validateFile(w http.ResponseWriter, r *http.Request, format string, filename string) bool {
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return false
	}

	file, _, err := r.FormFile(filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("file doesn't exist. %s", err.Error()), http.StatusBadRequest)
		return false
	}
	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read file. %s", err.Error()), http.StatusInternalServerError)
		return false
	}

	filetype := http.DetectContentType(buff)
	if filetype != format {
		{
			http.Error(w, "The provided file format '"+filetype+"' is not allowed. Please upload a "+format+" file",
				http.StatusBadRequest)
			return false
		}
	}
	return true
}

func saveFile(w http.ResponseWriter, r *http.Request, filename string, pathToSaveData string) bool {
	file, fileHeader, err := r.FormFile(filename)
	name := strings.TrimSuffix(filepath.Base(fileHeader.Filename), filepath.Ext(fileHeader.Filename))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving the file: %s", err.Error()), http.StatusBadRequest)
		return false
	}

	// Create the uploads folder if it doesn't already exist
	err = os.MkdirAll(pathToSaveData, os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("%s/%s%s", pathToSaveData, name, filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	return true
}

type Response struct {
	Message string `json:"message"`
}
