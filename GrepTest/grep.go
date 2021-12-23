package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func main() {
	//flag path
	path := flag.String("p", "", "File path to grep")

	//flag search
	search := flag.String("s", "", "Search string from the path file")

	//Define flag!
	flag.Parse()
	files, err := ioutil.ReadDir(*path) //pointer to use it on heap
	if err != nil {
		log.Fatal(err)
	}
	//for each files in directory
	for _, file := range files {
		if file.Mode().IsDir() {
			continue
		}
		// Read file inside directory
		data, err := ioutil.ReadFile(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		//details mime of files
		filext := http.DetectContentType(data)
		short := strings.Split(string(data), "\n")
		filecol := color.New(color.FgRed, color.Bold)
		linucol := color.New(color.FgYellow, color.Bold)
		linecol := color.New(color.FgWhite, color.Bold)
		linum := 1 //line start from 1
		for _, line := range short {
			// check logic if string in search var is exist in the file
			if strings.Contains(line, *search) {
				//for double check, checking extension
				if strings.Contains(filext, "text/plain") {
					if len(line) > 75 {
						line = line[:75]
						filecol.Printf("[%s]:\n", file.Name())
						linucol.Printf("\tLine-%d: ", linum)
						linecol.Printf("\n\t%s...\n", line)
					} else {
						filecol.Printf("[%s]:\n", file.Name())
						linucol.Printf("\tLine-%d: ", linum)
						linecol.Printf("\n\t%s\n", line)
					}
				} else {
					fmt.Println("Nil")
					break
				}
				fmt.Println("")
			}
			linum++
		}
	}
}
