package pathseparator

import (
	"fmt"
	"path"
)

func SplitPath() {
	var dir, file string

	dir, file = path.Split("css/main.css")
	// Another way would be
	// dir, file := path.Split("css/main.css")

	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
}
