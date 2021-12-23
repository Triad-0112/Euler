package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func compare(path string, path2 string) {
	var mapPath map[string]int
	mapPath = map[string]int{}
	tmpPath := path

	var mapPath2 map[string]int
	mapPath2 = map[string]int{}
	tmpPath2 := path2

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// fmt.Println(path, info.Size())

			// content, err := ioutil.ReadFile(path)
			// fmt.Println(string(content))

			// file := filepath.Base(path)
			// fmt.Println(file)
			trimPath := path[len(tmpPath):]
			mapPath[trimPath] = int(info.Size())
			return nil
		})

	err2 := filepath.Walk(path2,
		func(path2 string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			trimPath := path2[len(tmpPath2):]
			mapPath2[trimPath] = int(info.Size())
			return nil
		})

	for key, val := range mapPath {
		_, found := mapPath2[key]
		if found {
			if val == mapPath2[key] {
				//fmt.Println(key, "  \t:", val)
			} else {
				fmt.Println(tmpPath, key, "  \t **MODIFIED**")
				//fmt.Println(key, "  \t:", val, " **MODIFIED**")
			}
		} else {
			fmt.Println(tmpPath, key, "  \t **NEW**")
			//fmt.Println(key, "  \t:", val, " **NEW**")
		}

	}
	//fmt.Println("--------------")
	//for key, val := range mapPath2 {
	for key, _ := range mapPath2 {
		_, found := mapPath[key]
		if !found {
			//fmt.Println(key, "  \t:", val, " **DELETED**")
			fmt.Println(tmpPath2, key, "  \t: **DELETED**")
		}
		//fmt.Println(key, "  \t:", val)
	}

	if err != nil {
		log.Println(err)
	}
	if err2 != nil {
		log.Println(err2)
	}
}

func main() {
	// compare("E:\\cobalogy\\golang\\magicsoft\\compare", "")
	compare("..\\magicsoft\\compare\\example1", "..\\magicsoft\\compare\\example2")
	//fmt.Println("E:\\cobalogy\\golang\\magicsoft\\compare")
}
