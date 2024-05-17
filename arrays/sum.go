package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, numberSlice := range numbers {
		if len(numberSlice) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := numberSlice[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
