package internal

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode/utf8"
	"zlbenjamin/gofsmgnt/internal/utils"

	"github.com/gin-gonic/gin"
)

// Global variables

const dirMaxLenth = 248

const filesMaxSize = 10_000

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

	// Set gvdir
	gvdir = dir

	gviniting = true
	go startInitializeDir()

	c.JSON(http.StatusOK, gin.H{"status": 0, "message": "Start initing"})
}

// Start initialize dir
func startInitializeDir() {
	log.Println("Start initialize dir: ", gvdir)

	gvfiles = make([]string, 100)

	fsys := os.DirFS(gvdir)

	var total int
	fs.WalkDir(fsys, ".", func(path2 string, d fs.DirEntry, err error) error {
		if err != nil {
			gviniting = false

			// log.Fatal("err=", err)
			log.Println("err=", err)
			return fs.SkipDir
		}

		total++

		gvfiles = append(gvfiles, path2)
		if total > filesMaxSize {
			// Ignoring remaining
			log.Println("Skip more files in ", gvdir)
			return fs.SkipAll
		}

		return nil
	})

	log.Println("After initializing: ", gvdir, "total=", total)

	gviniting = false
}

// Query files by kw (key word, Case-sensitive)
// A maximum of 20 files are returned.
func QueryFiles(c *gin.Context) {
	if gviniting {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "The dir is initializing"})
		return
	}

	kw := c.Query("kw")
	kw = strings.Trim(kw, " ")

	if len(gvfiles) < 1 {
		c.JSON(http.StatusOK, gin.H{"kw": kw, "len": 0, "files": []string{}})
		return
	}

	if kw == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid kw, blank"})
		return
	}
	if kw == "/" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid kw, equal to /"})
		return
	}

	if utf8.RuneCountInString(kw) > 20 {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid kw, length exceeded"})
		return
	}

	maxLen := 20
	var files []string
	for _, name := range gvfiles {
		if strings.Contains(name, kw) {
			files = append(files, name)
			if len(files) >= maxLen {
				break
			}
		}
	}
	if len(files) < 1 {
		files = []string{}
	}

	c.JSON(200, gin.H{"kw": kw, "len": len(files), "files": files})
}
