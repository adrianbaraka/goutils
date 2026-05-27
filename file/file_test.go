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

func TestHumanReadableSize(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bytes  int64
		binary bool
		want   string
	}{
		// Decimal (Base 1000)
		{"0 bytes", 0, false, "0 B"},
		{"Only bytes", 800, false, "800 B"},
		{"Exactly 1 KB", 1000, false, "1 KB"},
		{"With many decimal Points", 1043532, false, "1.04 MB"},
		{"Rounding up (.567 -> .57)", 1567000, false, "1.57 MB"},
		{"Trailing zeros trimmed (.50 -> .5)", 1500000, false, "1.5 MB"},
		{"Large Gigabyte file", 4567000000, false, "4.57 GB"},
		{"Terabyte boundary", 1000000000000, false, "1 TB"},

		// Binary (Base 1024)
		{"0 bytes binary", 0, true, "0 B"},
		{"Just below KiB boundary", 1023, true, "1023 B"},
		{"Exactly 1 KiB", 1024, true, "1 KiB"},
		{"Slightly above 1 KiB", 1500, true, "1.46 KiB"},
		{"Exactly 1 MiB", 1048576, true, "1 MiB"},
		{"Typical 4GB Video File", 4294967296, true, "4 GiB"},
		{"Large scale TiB", 5497558138880, true, "5 TiB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := file.HumanReadableSize(tt.bytes, tt.binary)
			if got != tt.want {
				t.Errorf("HumanReadableSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
