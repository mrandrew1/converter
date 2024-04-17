package helpers

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const pathToCreateZip = "./to_download/"

func CreateZip(files []string, newFileName string) {
	// write code create a folder

	output, err := os.Create(fmt.Sprintf("%s%s.zip", pathToCreateZip, newFileName))
	if err != nil {
		panic(err)
	}
	defer output.Close()

	zipWriter := zip.NewWriter(output)
	defer zipWriter.Close()

	for _, file := range files {
		err = addFileToZip(zipWriter, file)
		if err != nil {
			panic(err)
		}
	}
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	file, err := os.Open("./results/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(filename)
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}
