/*
  For is Go's "while"
    - You can drop the semicolons.
    - C's "while" is spelled "for" in Go.
*/

package main
import "fmt"

func main() {
  sum := 1

  for sum < 100 {
    fmt.Println("Current Sum: ", sum)
    sum += sum
  }

  fmt.Println("Total Sum: ", sum)
}

