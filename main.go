package main

import (
	"awesomeProject1/modules/server"
)

const (
	pathToGetData    = "data.csv"
	pathToGetExample = "Заявление о выдаче приказа.docx"
)

func main() {
	server.Server()
}
