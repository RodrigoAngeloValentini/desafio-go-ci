
package main

import "fmt"

func Sum(x int, y int) int {
	return x + y;
}

func main() {
	res := Sum(5, 5)
  fmt.Println("5+5 = ", res)
}