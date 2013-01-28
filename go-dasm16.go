package main

import (
	"os"
)

var (
	Version = "0.0.1"
)

func main() {
	// check yo flags
	path := checkYoFlags()
	// read in file
	content, title, dir := readFile(path)
	// parse file
	content = parse(content)
	// write file
	writeFile(dir, content, title)
	os.Exit(0)
}

func checkYoFlags() (path string) {
	
	
	
	return
}