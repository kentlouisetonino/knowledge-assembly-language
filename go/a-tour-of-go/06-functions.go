/*
  Functions
    - A function can take zero or more arguments.

  Package References
    fmt   : https://pkg.go.dev/fmt
*/
package main
import "fmt"

func add(x int, y int) int {
  return x + y;
}

func main() {
  fmt.Println(add(10, 10));
}

