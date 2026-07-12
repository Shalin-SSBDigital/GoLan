package main
import "fmt"
func add (a int, b int) int {
	return a + b
}

func main4() {
	result := add(3, 4)
	fmt.Println("Result:", result)
}