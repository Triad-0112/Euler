package main

import (
	"fmt"
)

func main() {
	var jumlahdata int = 5
	var DataDiagram = make([]int, jumlahdata)
	DataDiagram = []int{2, 4, 1, 3, 0}
	fmt.Print("\n")
	Sorting(DataDiagram)
}
func Grafik(k int, l int, m []int) {
	for i := l; i >= 1; i-- { // membuat Data Diagram
		for j := 0; j < k; j++ {
			if m[j] >= i {
				fmt.Print(" | ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Print("\n")
	}
	for i := 0; i < k; i++ {
		fmt.Print("---")
	}
	fmt.Print("\n ")
	for i := 0; i < k; i++ {
		fmt.Print(m[i], "  ")
	}
	fmt.Print("\n")
}

func Sorting(m []int) {
	max := m[0]
	for _, value := range m { // Menemukan nilai maximum
		if value > max {
			max = value
		}
	}
	status := true
	for status {
		Grafik(len(m), max, m)
		status = false
		for i := 0; i < len(m)-1; i++ {
			if m[i] > m[i+1] {
				m[i], m[i+1] = m[i+1], m[i]
				status = true
			}
		}
	}
}
