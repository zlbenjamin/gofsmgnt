package src

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"
	"zlbenjamin/gofsmgnt/src/utils"

	"github.com/gin-gonic/gin"
)

// Global variables

const dirMaxLenth = 248

// Current dir
var gvdir string

// file paths under gvdir, directories and files.
var gvfiles []string

// true: initing, stop new initing.
var gviniting bool

// Initialize based on a dir
// query parameter: dir, absolute path
func InitFs(c *gin.Context) {
	dir := c.Query("dir")
	dir = strings.Trim(dir, " ")

	// Check dir
	if dir == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid dir, blank"})
		return
	}
	if utf8.RuneCountInString(dir) > dirMaxLenth {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid dir, length exceeded"})
		return
	}

	if !utils.CheckFileExists(dir) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "No such dir"})
		return
	}

	if !utils.CheckFileExists(dir) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "No such dir"})
		return
	}

	if gviniting {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "The last one is initializing"})
		return
	}

	gviniting = true
	go startInitializeDir()

	c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "Start initing"})
}

// Start initialize dir
func startInitializeDir() {
	fmt.Println("Start initialize dir: ", gvdir)

	// TODO

	gviniting = false
}
