package sieve

import "math"

// Sieve returns a list of all prime numbers up to limit
func Sieve(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false

	for i := 2; i <= int(math.Sqrt(float64(limit))); i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	var result []int
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			result = append(result, i)
		}
	}
	return result
}
