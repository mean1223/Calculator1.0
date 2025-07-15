package calculator

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	a := 10
	var b int = 5
	fmt.Println(Add(a, b))
}
