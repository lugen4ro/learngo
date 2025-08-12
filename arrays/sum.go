package main

func Sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func SumAll(arrsToSum ...[]int) []int {

	lengthOfNumber := len(arrsToSum)
	// var sums []int
	sums := make([]int, lengthOfNumber)

	for i, arr := range arrsToSum {
		sums[i] = Sum(arr)
	}

	return sums
}
