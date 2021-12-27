package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Search searches files for lines matching a regular expression pattern.
func Search(pattern string, flags, files []string) []string {
	var lineNums, namesOnly, ignoreCase, invert, entireLines bool
	for _, flag := range flags {
		switch flag {
		case "-n":
			lineNums = true
		case "-l":
			namesOnly = true
		case "-i":
			ignoreCase = true
		case "-v":
			invert = true
		case "-x":
			entireLines = true
		}
	}
	if ignoreCase {
		pattern = strings.ToLower(pattern)
	}
	if entireLines {
		pattern = "^" + pattern + "$"
	}
	rx := regexp.MustCompile(pattern)
	matches := []string{}
	for _, fname := range files {
		file, _ := os.Open(fname)
		scanner := bufio.NewScanner(file)
		linenum := 1
		for scanner.Scan() {
			line := scanner.Text()
			var match bool
			if ignoreCase {
				match = rx.MatchString(strings.ToLower(line))
			} else {
				match = rx.MatchString(line)
			}
			if invert {
				match = !match
			}
			if match {
				var output string
				if namesOnly {
					matches = append(matches, fname)
					break
				}
				if len(files) > 1 {
					output = fname + ":"
				}
				if lineNums {
					output += fmt.Sprintf("%d:", linenum)
				}
				output += line
				matches = append(matches, output)
			}
			linenum++
		}
		file.Close()
	}
	return matches
}
func main() {
	a := "write"
	files := []string{"teststory.txt"}
	flags := []string{"-l"}
	Search(a, flags, files)

}
