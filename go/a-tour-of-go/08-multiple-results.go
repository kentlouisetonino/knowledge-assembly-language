/*
  Multiple Results
    - A function can return any number of results.
    - The "swap" function returns two strings.

  Package References
    fmt   : https://pkg.go.dev/fmt
*/

package main
import "fmt"

func swap(str1, str2 string) (string, string) {
  return str1, str2;
}

func main() {
  a, b := swap("Cruel", "World!");
  fmt.Println(a, b)
}

