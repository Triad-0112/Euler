package main

import (
	"flag"
	"fmt"
)

func main() {
	anag1 := flag.String("1", "", "Susunan kata yang akan dicek")
	anag2 := flag.String("2", "", "Susunan kata yang akan dibandingkan")
	flag.Parse()
	checkForAnagrams(*anag1, *anag2)
}

func checkForAnagrams(anag1, anag2 string) {
	if len(anag1) != len(anag2) {
		fmt.Printf("%s and %s are not anagrams", anag1, anag2)
		return
	}

	for i := 0; i < len(anag1); i++ {
		if anag1[i] != anag2[len(anag2)-i-1] {
			fmt.Printf("%s and %s are not anagrams", anag1, anag2)
			return
		}
	}
	fmt.Printf("%s and %s are anagrams", anag1, anag2)
}
