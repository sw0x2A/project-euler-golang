// 2520 is the smallest number that can be divided by each of the numbers
// from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of
// the numbers from 1 to 20?

package main

func p5(k uint) uint {
	p := []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	r := uint(1)
	var d uint
	for i := range p {
		d = p[i]
		for d <= k {
			d *= p[i]
		}
		d /= p[i]
		r *= d
	}
	return r
}
