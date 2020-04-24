// Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
// Example 1:
// Input: [5,7]
// Output: 4
// Example 2:
// Input: [0,1]
// Output: 0

// 0000 0
// 0001 1
// 0010 2
// 0011 3
// 0100 4
// 0101 5
// 0110 6
// 0111 7
// 1000 8
// 1001 9
// 1010 10
// 1011 11
// 1100 12
// 1101 13
// 1110 14
// 1111 15
// 10000 16
// 10001 17
// 10010 18
// 10011 19
// 10100 20
// 10101 21
// 10110 22
// 10111
// 11000
// 11001
// 11010
// 11011
// 11100
// 11101
// 11110
// 11111

package main

import (
	"math"
)

// the best solution
func rangeBitwiseAnd(m int, n int) int {
	for n > m {
		n = n & (n - 1)
	}
	return n & m
}

// my solution
func my_rangeBitwiseAnd(m int, n int) int {
	var sumBitwise int

	if n <= m+1 {
		return m & n
	}
	for bit := 15; bit > int(math.Log2(float64(n-m))); bit-- {
		bitCompare := int(math.Pow(2, float64(bit)))
		if m&bitCompare > 0 && n&bitCompare > 0 {
			sumBitwise += bitCompare
		}
	}
	return sumBitwise
}

// Brute force
func _rangeBitwiseAnd(m int, n int) int {
	var bitwiseAnd int = m

	for i := m + 1; i <= n; i++ {
		bitwiseAnd &= i
	}

	return bitwiseAnd
}

// case 8 not pass
func ___rangeBitwiseAnd(m int, n int) int {
	var bit int
	if m == n {
		return m
	}
	if n == m+1 {
		return n & m
	}
	for ; bit < 16; bit++ {
		if m>>1 == 0 {
			break
		}
		m >>= 1
		n >>= 1
	}
	if m != n {
		return 0
	}
	return int(math.Pow(float64((m&n)*2), float64(bit)))
}
