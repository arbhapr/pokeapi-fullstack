package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomCatchSuccess() bool {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(2) == 0 // 50% probability
}

// GenerateFibonacci returns the nth Fibonacci number
func GenerateFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// GenerateRandomNickname creates a random nickname
func GenerateRandomNickname(pikachuName string) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	adjectives := []string{"Brave", "Mighty", "Swift", "Shadow", "Lucky"}
	return fmt.Sprintf("%s %s", adjectives[r.Intn(len(adjectives))], pikachuName)
}

// RandomReleaseSuccess generates a random number and returns true if it's prime, otherwise false
func RandomReleaseSuccess() bool {
	// Generate a random number between 1 and 100 (or any range you prefer)
	randomNumber := rand.Intn(100) + 1

	// Check if the number is prime
	return isPrime(randomNumber)
}

// isPrime checks if a given number is a prime number
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}
