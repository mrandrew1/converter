package createWordFiles

func GenerateWordFilesFromCsV(pathToGetData string, pathToGetExample string, newFileName string) {
	csvDataArray, columns := getDataFromCsv(pathToGetData)
	createWordFiles(columns, csvDataArray, pathToGetExample, newFileName)
}
