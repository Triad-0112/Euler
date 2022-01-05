package main

import (
	"fmt"
	"os"
)

func main() {
	//Get Current Path
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("starting dir:", wd)
	if err := os.Chdir("/"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := os.Getwd(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Final dir:", wd)
}
