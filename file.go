package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func readFile(path string) (content string, title string, dir string) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		os.Exit(1)
	} else {
		content = string(f)
		title = path[strings.LastIndex(path, "/")+1:]
		dir = path[0:strings.LastIndex(path, title)] 
		println("title: "+title)
		println("dir: "+dir)
		println("path: "+path)
	}
	return
}

func writeFile(dir string, content string, title string) {
	err := ioutil.WriteFile(dir+title+".dasm16", []byte(content), 0664)
	if err != nil {
		os.Exit(1)
	}
}