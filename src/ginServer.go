package src

import (
	"fmt"
	"io"
	"log"
	"os"

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
	addr := fmt.Sprintf("localhost:%d", gport)
	log.Println("Start server: ", addr)
	router.Run(addr)
}

// Configure router
func configRouter(router *gin.Engine) {
	router.POST("/api/fs/init", InitFs)
	router.GET("/api/fs/query", QueryFiles)

	router.GET("/api/sys/meminfo", SysMemInfo)
}
