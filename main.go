package main

import "fmt"

func main() {

	var testInput = []int{3, 0, -2, 4, -6}

	resultReversearray := Reversearray(testInput)

	k := -4

	resultisTailSum := isTailSum(resultReversearray, k)

	fmt.Println(resultisTailSum)

}

func isTailSum(testInput []int, k int) int {

	n := 0
	result := 0

	if len(testInput) > 0 {

		for i := 0; i < len(testInput); i++ {

			result += testInput[i]

			if result == k {

				n = i + 1
				break

			} else {

				n = 0
			}

		}

	} else {
		n = 0
	}
	return n

}

func Reversearray(array []int) []int {

	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}

	return array
}
