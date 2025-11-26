package main
import "fmt"

func sortArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func drawDiamond(n int) {
	for i := 1; i <= n; i++ {
		for j := i; j < n; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := n - 1; i >= 1; i-- {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func main(){
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	sortedArr := sortArray(arr)
	fmt.Println("Sorted array:", sortedArr)
	drawDiamond(10)
}