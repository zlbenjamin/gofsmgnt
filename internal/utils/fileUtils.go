package utils

import "os"

// Check if the path exists
// path: absulate path
// return true, exists.
func CheckFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
