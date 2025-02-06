package utils

import "math"


// IsPrime determines if a given integer is a prime number.
// A prime number is greater than 1 and divisible only by 1 and itself.
func IsPrime(n int) bool{
	if n <= 1 {
		// Numbers less than or equal to 1 are not prime
		return false
	}
	// Check divisors from 2 to the square root of n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0{
			return false
		}

	}
	return true 
}

// IsPerfect determines if a given integer is a perfect number.
// A perfect number is a positive integer that equals the sum of its proper divisors (excluding itself).
func IsPerfect(n int) bool{
	sum := 0

	// Iterate through all numbers less than n
	for i := 1; i < n; i++ {
		if n % i == 0 {
			sum += i
		}
	}
	// Check if the sum of divisors equals the original number
	return sum == n
}

// IsArmstrong determines if a given integer is an Armstrong number.
// An Armstrong number (or narcissistic number) is a number equal to the sum of its digits raised to the power of 3.
func IsArmstrong(n int) bool {
	sum, temp := 0, n

	// Extract and process each digit of the number
	for temp > 0 {
		digit := temp % 10                       // Get the last digit
		sum += int(math.Pow(float64(digit),3))  // Add the cube of the digit to the sum
		temp /= 10                              // Remove the last digit
	}

	// Check if the sum of the cubes equals the original number
	return sum == n
}


// DigitSum calculates the sum of the digits of a given integer.
func DigitSum(n int) int{
	sum := 0
	for n > 0 {
		sum += n % 10   // Add the last digit
		n /=10          // Remove the last digit
	}

	return sum
}
// Parity determines whether a given integer is odd or even.
// Returns "even" if the number is divisible by 2, otherwise "odd"

func Parity(n int) string{
	if n % 2 == 0 {

		return "even"
	} 

	return "odd"
}