/*
  Every go program is made of packages. Program start running
  in package "main".

  By convention, the package name is the same as the last element
  of the import path. For instance, the "math/rand" package
  comprises files that begin with the statement "package rand".

  Package References:
    fmt         : https://pkg.go.dev/fmt
    math/rand   : https://pkg.go.dev/math/rand
*/

package main
import "fmt"
import "math/rand"

func main() {
  fmt.Println("The number is", rand.Intn(10));
}

