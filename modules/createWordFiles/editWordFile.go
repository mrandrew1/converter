package createWordFiles

import (
	"fmt"
	"github.com/nguyenthenguyen/docx"
)

func createWordFiles(columns []string, data [][]string, path string, newFileName string) {
	pathToSave := "./results/"
	pathToGetExamples := "./uploads/examples/"
	// Read from docx file
	r, err := docx.ReadDocxFile(pathToGetExamples + path) //  example path

	if err != nil {
		panic(err)
	}

	for i, file := range data {
		docx1 := r.Editable()
		for i, _ := range columns {
			docx1.ReplaceRaw(columns[i], file[i], -1)
		}
		docx1.WriteToFile(fmt.Sprintf("%s%s_%d.docx", pathToSave, newFileName, i+1))
	}

	r.Close()

}
