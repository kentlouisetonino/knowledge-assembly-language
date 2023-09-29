/*
  For Continued
    - The init and post statements are optional.
*/

package main
import "fmt"

func main() {
  sum := 1

  for ; sum < 100; {
    fmt.Println("Current Sum: ", sum)
    sum += sum
  }

  fmt.Println("Total Sum: ", sum)
}

