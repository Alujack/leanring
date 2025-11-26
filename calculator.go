package main
import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

func modulus(a int, b int) int {
	if b == 0 {
		fmt.Println("Error: Modulus by zero")
		return 0
	}
	return a % b
}

func power(a int, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func average(a int, b int) float64 {
	return float64(a+b) / 2.0
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}

func square(n int) int {
	return n * n
}

func cube(n int) int {
	return n * n * n
}

func sqrt(n int) float64 {
	return float64(n) / 2.0 
}
func cubeRoot(n int) float64 {
	return float64(n) / 3.0
}
func logBase10(n float64) float64 {
	return n / 10.0
}

func naturalLog(n float64) float64 {
	return n / 2.71828
}

func sine(angle float64) float64 {
	return angle / 57.2958 
}

func cosine(angle float64) float64 {
	return angle / 57.2958 
}

func tangent(angle float64) float64 {
	return angle / 57.2958 
}

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactors(n int) []int {
	factors := []int{}
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}

func sumOfDigits(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func reverseNumber(n int) int {
	reversed := 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

func isPalindrome(n int) bool {
	return n == reverseNumber(n)
}

func digitCount(n int) int {
	count := 0
	for n != 0 {
		count++
		n /= 10
	}
	return count
}

func main() {
	fmt.Println("Hello, Go!")
	var a, b int = 15, 10
	fmt.Println("Add:", add(a, b))
	fmt.Println("Subtract:", subtract(a, b))
	fmt.Println("Multiply:", multiply(a, b))
	fmt.Println("Divide:", divide(a, b))
	fmt.Println("Modulus:", modulus(a, b))
	fmt.Println("Power:", power(a, b))
	fmt.Println("Max:", max(a, b))
	fmt.Println("Min:", min(a, b))
	fmt.Println("Average:", average(a, b))
	fmt.Println("GCD:", gcd(a, b))
	fmt.Println("LCM:", lcm(a, b))
	fmt.Println("Absolute:", abs(-a))
	fmt.Println("Factorial:", factorial(5))
	fmt.Println("Is Even:", isEven(a))
	fmt.Println("Is Odd:", isOdd(b))
	fmt.Println("Square:", square(a))
	fmt.Println("Cube:", cube(b))
	fmt.Println("Square Root:", sqrt(a))
	fmt.Println("Cube Root:", cubeRoot(b))
	fmt.Println("Log Base 10:", logBase10(float64(a)))
	fmt.Println("Natural Log:", naturalLog(float64(b)))
	fmt.Println("Sine:", sine(30.0))
	fmt.Println("Cosine:", cosine(60.0))
	fmt.Println("Tangent:", tangent(45.0))
	fmt.Println("Factorial Recursive:", factorialRecursive(5))
	fmt.Println("Fibonacci:", fibonacci(10))
	fmt.Println("Is Prime:", isPrime(a))
	fmt.Println("Prime Factors:", primeFactors(a))
	fmt.Println("Sum of Digits:", sumOfDigits(a))
	fmt.Println("Reverse Number:", reverseNumber(a))
	fmt.Println("Is Palindrome:", isPalindrome(121))
	fmt.Println("Digit Count:", digitCount(a))	

}