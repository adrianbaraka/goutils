package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

// HumanReadableSize takes given bytes number and converts to a human readable string.
// The binary determines if KiB (1024) or decimal KB (1000) will be used.
func HumanReadableSize(bytes int64, binary bool) string {
	divisor := 1000.0
	i := ""
	if binary {
		divisor = 1024.0
		i = "i"
	}

	if bytes < int64(divisor) {
		return fmt.Sprintf("%v B", bytes)
	}

	symbols := []string{"", "K", "M", "G", "T", "P", "E"}

	b := float64(bytes)
	rounds := 0
	for b >= divisor && rounds < len(symbols)-1 {
		b /= divisor
		rounds++
	}


	valueStr := strconv.FormatFloat(b, 'f', 2, 64)

	// Clean up trailing zeros
	valueStr = strings.TrimSuffix(valueStr, ".00")
	if strings.Contains(valueStr, ".") {
		valueStr = strings.TrimSuffix(valueStr, "0")
	}

	return fmt.Sprintf("%s %s%sB", valueStr, symbols[rounds], i)
}
