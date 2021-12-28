package main

import (
	"fmt"
)

var retAry [][]int

func main() {
	fmt.Println(HowSum(5, []int{1, 2, 5}))
}

func HowSum(targetNum int, numbers []int) []int {
	retAry = make([][]int, targetNum+1)
	retAry[0] = make([]int, 0)
	for i := 0; i <= targetNum; i++ {
		if retAry[i] != nil {
			for _, num := range numbers {
				if i+num <= targetNum {
					fmt.Print("Before:", i, " round, num =", num, retAry, "\n")
					// 這樣做，在使用Append修改retAry[5]的時候，retAry[4]的內容也會被改掉
					// Before:3 round, num =2 [[] [1] [1 1] [1 1 1] [1 1 1 1] [5]]
					// After :3 round, num =2 [[] [1] [1 1] [1 1 1] [1 1 1 2] [1 1 1 2]]
					retAry[i+num] = append(retAry[i], num)
					// 下面這個就不會
					// Before:3 round, num =2 [[] [1] [1 1] [1 1 1] [1 1 1 1] [5]]
					// After :3 round, num =2 [[] [1] [1 1] [1 1 1] [1 1 1 1] [2 1 1 1]]
					//retAry[i+num] = append([]int{num}, retAry[i]...)

					fmt.Print("After :", i, " round, num =", num, retAry, "\n\n")

				}
			}
		}
	}
	return retAry[targetNum]
}
