package array_slice

func Sum(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		var tail []int
		if len(numbers) > 0 {
			tail = numbers[1:]
		}
		sums = append(sums, Sum(tail))
	}
	return sums
}
