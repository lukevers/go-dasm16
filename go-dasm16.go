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
	switch len(os.Args) {
		default:
			println("\033[31mError\033[0m: Too many arguments! Stopping.")
			os.Exit(0)
		case 1:
			println("\033[31mError\033[0m: No file included! Stopping.")
			os.Exit(0)
		case 2:
			path = os.Args[1]
			if _, err := os.Stat(path); os.IsNotExist(err) {
				println("\033[31mError\033[0m: File does not exist! Stopping.")
				os.Exit(0)
			}
	}

	return
}