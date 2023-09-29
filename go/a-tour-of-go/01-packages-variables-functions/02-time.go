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

  Time (Package)
    Wall Clock
      - Subject to changes for clock synchronization.
      - The general rule is that the wall clock is for telling time.

    Monotonic Clock
      - Is not subject to change.
      - Is for measuring time.
      - A time source that won't ever jump forward or backward (due to Network Time
        Protocol or Daylight Saving Time updates).
*/

package main
import "fmt"
import "time"

func main() {
  fmt.Println("Current Time: ", time.Now());
}

