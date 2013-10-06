// https://projecteuler.net/problem=11

package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func p11(file string) int {
	fileBuf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fileStr := string(fileBuf)
	lines := strings.SplitN(fileStr, "\n", 20)
	g := make([][]int, 20)
	var columns []string
	for i, v := range lines {
		g[i] = make([]int, 20)
		columns = strings.SplitN(v, " ", 20)
		for j := range g[i] {
			g[i][j], _ = strconv.Atoi(columns[j])
		}
	}

	l, t, i, j := 0, 0, 0, 0

	// top-left to bottom-right
	for i = 0; i <= 16; i++ {
		for j = 0; j <= 16; j++ {
			t = g[i][j] * g[i+1][j+1] * g[i+2][j+2] * g[i+3][j+3]
			if l < t {
				l = t
			}
		}
	}

	// top-right to bottom-left
	for i = 0; i <= 16; i++ {
		for j = 3; j < 20; j++ {
			t = g[i][j] * g[i+1][j-1] * g[i+2][j-2] * g[i+3][j-3]
			if l < t {
				l = t
			}
		}
	}

	// left
	for i = 0; i <= 16; i++ {
		for j = 0; j < 20; j++ {
			t = g[i][j] * g[i+1][j] * g[i+2][j] * g[i+3][j]
			if l < t {
				l = t
			}
		}
	}

	// down
	for i = 0; i < 20; i++ {
		for j = 0; j <= 16; j++ {
			t = g[i][j] * g[i][j+1] * g[i][j+2] * g[i][j+3]
			if l < t {
				l = t
			}
		}
	}

	return l
}
