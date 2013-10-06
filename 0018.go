// https://projecteuler.net/problem=18

package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func p18() int {
	fileBuf, err := ioutil.ReadFile("0018_input.txt")
	if err != nil {
		panic(err)
	}
	fileStr := string(fileBuf)
	lines := strings.SplitN(fileStr, "\n", -1)
	var line []string
	arr := make([][]int, 15)
	for i := range arr {
		line = strings.SplitN(lines[i], " ", -1)
		arr[i] = make([]int, len(line))
		for j := range line {
			arr[i][j], _ = strconv.Atoi(line[j])
		}
	}
	for j := 14; j >= 0; j-- {
		for i := 0; i < j; i++ {
			if arr[j][i] > arr[j][i+1] {
				arr[j-1][i] += arr[j][i]
			} else {
				arr[j-1][i] += arr[j][i+1]
			}
		}
	}
	return arr[0][0]
}
