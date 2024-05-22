package main

import (
	"log"
	"zlbenjamin/gofsmgnt/src"
	"zlbenjamin/gofsmgnt/src/utils"
)

// the main()
func main() {
	log.Println("System info: ", utils.GetSystemInfo())

	src.StartServer()
}
