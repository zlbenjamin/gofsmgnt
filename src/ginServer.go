package src

import (
	"fmt"
	"io"
	"log"
	"os"
	"zlbenjamin/gofsmgnt/src/utils"

	"github.com/gin-gonic/gin"
)

const gport = 40000

// Start Web Server
func StartServer() {
	// logging to file
	// a windows file
	logFile := "d:\\run\\golog\\2024.log"
	f, _ := os.Create(logFile)
	// #1
	// gin.DefaultWriter = io.MultiWriter(f)
	// #2
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	configRouter(router)

	// router.Run()
	// or
	// addr := fmt.Sprintf("localhost:%d", gport) // no: in container
	// Listen and serve on 0.0.0.0:gport
	addr := fmt.Sprintf(":%d", gport) // yes
	log.Println("Start server: ", addr)
	log.Println("Memory: ", utils.RuntimeMem())
	router.Run(addr)
}

// Configure router
func configRouter(router *gin.Engine) {
	router.POST("/api/fs/init", InitFs)
	router.GET("/api/fs/query", QueryFiles)

	router.GET("/api/sys/meminfo", SysMemInfo)
	router.GET("/api/sys/info", SysInfo)
}
