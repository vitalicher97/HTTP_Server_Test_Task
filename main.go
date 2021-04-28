package main

import (
	"./controller"
	"./storage"
)

func main() {

	storage.InitLitMemory()
	controller.MainController()

}
