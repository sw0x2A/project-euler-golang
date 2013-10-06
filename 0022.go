// https://projecteuler.net/problem=22

package main

import (
	"io/ioutil"
	"sort"
	"strings"
)

func p22() int {
	fileBuf, err := ioutil.ReadFile("0022_input.txt")
	if err != nil {
		panic(err)
	}
	fileStr := string(fileBuf)
	names := strings.SplitN(fileStr, ",", -1)
	for i := range names {
		names[i] = strings.SplitN(names[i], "\"", -1)[1]
	}
	sort.Strings(names)
	sum := 0
	for i, v := range names {
		sum += alphabeticalValue(v) * (i + 1)
	}
	return sum
}

func alphabeticalValue(s string) int {
	sum := 0
	for _, v := range s {
		sum += int(v) - 64
	}
	return sum
}
