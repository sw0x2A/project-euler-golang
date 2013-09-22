// Work out the first ten digits of the sum of the following one-hundred
// 50-digit numbers.

package main

import (
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

func p13() string {
	fileBuf, err := ioutil.ReadFile("0013_input.txt")
	if err != nil {
		panic(err)
	}
	fileStr := string(fileBuf)
	lines := strings.SplitN(fileStr, "\n", 100)
	result := ""
	carryout, digit := 0, 0
	for i := 49; i >= 0; i-- {
		sum := carryout
		for j := range lines {
			num, _ := strconv.Atoi(string(lines[j][i]))
			sum += num
		}
		digit = sum % 10
		carryout = sum / 10
		result = strconv.Itoa(digit) + result
	}
	result = strconv.Itoa(carryout) + result
	return result[0:10]
}

func p13big() string {
	fileBuf, err := ioutil.ReadFile("0013_input.txt")
	if err != nil {
		panic(err)
	}
	fileStr := string(fileBuf)
	s := strings.Split(fileStr, "\n")
	n := new(big.Int).SetInt64(0)
	for _, v := range s {
		if t, ok := new(big.Int).SetString(v, 10); ok {
			n.Add(n, t)
		}
	}
	return n.String()[0:10]
}
