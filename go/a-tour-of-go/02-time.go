/*
  Package References:
  fmt   : https://pkg.go.dev/fmt
  time  : https://pkg.go.dev/time

  Note:
    - The time calculations is always based on the Gregorian Calendarian. Which
      is a solar calendar with 12 months of 28-31 days each. The year in both
      in calendars consists of 365 days, with a leap day being added to February
      in the leap years.
    - In this package the calculations has no leap seconds.
*/

package main
import "fmt"
import "time"

func main() {
  fmt.Println("Fucking time: ", time.Now());
}

