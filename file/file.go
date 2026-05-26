package file

import (
	"os"
	"path/filepath"
	"strings"
)

// Returns the given filename stripped of the last extension eg /home/document.pdf /home/document
func FileNameNoExtension(filename string) string {
	if filename == "" {
		return ""
	}
	// handle trailing slash
	if strings.HasSuffix(filename, string(os.PathSeparator)) {
		return filename
	}

	base := filepath.Base(filename)

	ext := filepath.Ext(base)

	// handle hidden files eg .gitignore
	if base == ext {
		ext = ""
	}

	parent := filepath.Dir(filename)

	//	fmt.Println(base)
	val := strings.TrimSuffix(base, ext)

	return filepath.Join(parent, val)
}
