package src

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const gport = 40000

// Start Web Server
func StartServer() {
	router := gin.Default()

	configRouter(router)

	// router.Run()
	// or
	addr := fmt.Sprintf("localhost:%d", gport)
	router.Run(addr)
}

// Configure router
func configRouter(router *gin.Engine) {
	router.POST("/api/fs/init", InitFs)
	router.GET("/api/fs/query", QueryFiles)
}
