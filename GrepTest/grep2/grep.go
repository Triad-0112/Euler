package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func OneFileSearch(flag map[rune]bool, filename string, multifile bool, re *regexp.Regexp) []string {
	results := []string{}
	file, err := os.Open(filename)
	if err != nil {
		panic("failed to open")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	n := 0
	for scanner.Scan() {
		n++
		line := scanner.Text()
		find := re.MatchString(line)
		if (find && !(flag['v'])) || (!(find) && flag['v']) {
			if flag['l'] {
				fmt.Println(filename)
				return []string{filename}
			}
			if flag['n'] {
				line = fmt.Sprintf("%d:%s", n, line)
			}
			if multifile {
				line = fmt.Sprintf("%s:%s", filename, line)
			}
			results = append(results, line)
		}
	}
	return results
}
func Search(pattern string, flags, files []string) []string {
	results := []string{}
	// scan flags
	flag := map[rune]bool{
		'i': false,
		'n': false,
		'l': false,
		'v': false,
		'x': false,
	}
	for _, f := range flags {
		switch f {
		case "-i":
			flag['i'] = true
			pattern = "(?i)" + pattern
		case "-n":
			flag['n'] = true
		case "-l":
			flag['l'] = true
		case "-v":
			flag['v'] = true
		case "-x":
			flag['x'] = true
			pattern = "^" + pattern + "$"
		}
	}
	for _, filename := range files {
		resonefile := OneFileSearch(flag, filename, len(files) > 1, regexp.MustCompile(pattern))
		results = append(results, resonefile...)
	}
	return results
}
func main() {
}
