package main

import (
	"fmt"
	"strings"

	"github.com/adrianbaraka/goutils/file"
)

func main() {
	//cli.RunCmd(true, false, true, "ls", "-la")

	strings.Contains("Hello", "h")

	var hi bool

	if hi {
		fmt.Println("Hi")
	}

	fmt.Println(file.HumanReadableSize(1000, false))

}
