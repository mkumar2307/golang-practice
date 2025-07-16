/* Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so */

package main

import (
    "fmt"
)

// GenDisplaceFn returns a function that calculates displacement based on time t
func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
    return func(t float64) float64 {
        return 0.5*a*t*t + vo*t + so
    }
}

func main() {
    var a, vo, so, t float64

    // Prompt user for input
    fmt.Print("Enter acceleration (a): ")
    fmt.Scan(&a)

    fmt.Print("Enter initial velocity (vo): ")
    fmt.Scan(&vo)

    fmt.Print("Enter initial displacement (so): ")
    fmt.Scan(&so)

    // Generate the displacement function
    fn := GenDisplaceFn(a, vo, so)

    // Prompt for time
    fmt.Print("Enter time (t): ")
    fmt.Scan(&t)

    // Compute and display displacement
    displacement := fn(t)
    fmt.Printf("Displacement after %.2f seconds is %.2f\n", t, displacement)
}