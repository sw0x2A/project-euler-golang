// https://projecteuler.net/problem=26

package main

func p26(n int) int {
	sequenceLength := 0

	for i := n; i > 1; i-- {
		if sequenceLength >= i {
			break
		}

		foundRemainders := make([]int, i)
		value, position := 1, 0

		for foundRemainders[value] == 0 && value != 0 {
			foundRemainders[value] = position
			value *= 10
			value %= i
			position++
		}

		if position-foundRemainders[value] > sequenceLength {
			sequenceLength = position - foundRemainders[value]
		}
	}
	return sequenceLength + 1
}
