package main

import "fmt"

/*
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.



Example 1:

Input: x = 121
Output: true
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
Example 4:

Input: x = -101
Output: false


Constraints:

-231 <= x <= 231 - 1

Follow up: Could you solve it without converting the integer to a string?
*/

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var revertedNum int
	for x > revertedNum {
		revertedNum = revertedNum*10 + x%10
		x /= 10
	}
	return x == revertedNum || x == revertedNum/10
}

func main() {
	fmt.Printf("x = 121, %t\n", isPalindrome(121))
	fmt.Printf("x = -121, %t\n", isPalindrome(-121))
	fmt.Printf("x = 10, %t\n", isPalindrome(10))
	fmt.Printf("x = -101, %t\n", isPalindrome(-101))
}