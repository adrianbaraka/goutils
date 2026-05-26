package file_test

import (
	"testing"

	"github.com/adrianbaraka/goutils/file"
)

func TestFileNameNoExtension(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		filename string
		want     string
	}{
		// TODO: Add test cases.
		{"Windows", `C:\Users\abc\document.pdf`, `C:\Users\abc\document`},
		{name: "empty"},
		{"Multi dot", "file.name.tar.gz", "file.name.tar"},
		{"No extension", "/etc/passwd", "/etc/passwd"},
		{"Root Path", "/", "/"},
		{"Hidden file", "/home/user/.config", "/home/user/.config"},
		{"Double hidden file", "/home/.render.yaml", "/home/.render"},
		{"Trailing slash", "/home/", "/home/"},
		{"Trailing slash windows", `C:\Users\`, `C:\Users\`},
		{"Only extension", ".mkv", ".mkv"},
		{"Relative Parent", "../../video.mkv", "../../video"},
		{"Current Dir", "./video.mkv", "video"},
		{"spaces", " file new .mkv", " file new "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := file.FileNameNoExtension(tt.filename)
			if got != tt.want {
				t.Errorf("FileNameNoExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
