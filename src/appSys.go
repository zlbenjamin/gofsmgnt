package src

import (
	"net/http"
	"zlbenjamin/gofsmgnt/src/utils"

	"github.com/gin-gonic/gin"
)

// Get memery info
func SysMemInfo(c *gin.Context) {
	mi := utils.RuntimeMem()
	c.JSON(http.StatusOK, gin.H{"value": mi})
}
