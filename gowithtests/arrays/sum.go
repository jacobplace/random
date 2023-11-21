package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var retVal []int
	for _, i := range numbersToSum {
		retVal = append(retVal, Sum(i))
	}
	return retVal
}
