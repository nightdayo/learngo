package main
import (
"fmt"
"stacker"
)
func main() {
	var haystack stacker.Stack
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin","clip","needle"})
	haystack.Push(81.52)
	var a int = 0
	for {
		a++
		item,err := haystack.Pop()
		fmt.Println(item," + ",a)
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
