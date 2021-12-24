package diffsquares

func SquareOfSum(n int) int {
	firstN := 0
	for i := 1; i <= n; i++ {
		firstN += i
	}
	firstN = firstN * firstN
	return firstN
}

func SumOfSquares(n int) int {
	firstN := 0
	for i := 1; i <= n; i++ {
		firstN = firstN + (i * i)
	}
	return firstN
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
