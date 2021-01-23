package main

import (
	"fmt"
	"math"
)

/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).



Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21
Example 4:

Input: x = 0
Output: 0


Constraints:

-231 <= x <= 231 - 1
*/

func reverse(x int) int {
	var rev int
	for x != 0 {
		pop := x % 10
		x /= 10
		rev = rev*10 + pop
	}
	if rev < math.MinInt32 || rev > math.MaxInt32 {
		return 0
	}
	return rev
}

func main() {
	fmt.Printf("x = 123, ans = %d\n", reverse(123))
	fmt.Printf("x = -123, ans = %d\n", reverse(-123))
	fmt.Printf("x = 120, ans = %d\n", reverse(120))
	fmt.Printf("x = 0, ans = %d\n", reverse(0))
}
