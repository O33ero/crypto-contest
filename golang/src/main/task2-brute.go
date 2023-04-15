package main

import (
	. "golang/src/util"
	"log"
)

func main() {
	text, err := GetFileContent(task2File)
	if err != nil {
		return
	}

	log.Println(text)
}
