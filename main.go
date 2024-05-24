package main

import (
	"log"
	internal "zlbenjamin/gofsmgnt/internal"
	"zlbenjamin/gofsmgnt/internal/utils"
)

// the main()
func main() {
	log.Println("System info: ", utils.GetSystemInfo())

	utils.GetDiskInfo("c:")

	internal.StartServer()
}
