/*package thefarm

// See types.go for the types defined for this exercise.
import (
	"errors"
	"fmt"
)

// TODO: Define the SillyNephewError type here.
// DivideFood computes the fodder amount per cow for the given cows.
// TODO: Define the SillyNephewError type here.
// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	value, err := weightFodder.FodderAmount()
	if err == nil && value >= 0 && cows > 0 {
		return value / float64(cows), nil
	}
	if err == ErrScaleMalfunction && value > 0 && cows > 0 {
		return value / float64(cows) * 2, nil
	}
	if value < 0 {
		return 0, errors.New("Negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("Division by zero")
	}
	if cows < 0 {
		return 0, errors.New(fmt.Sprintf(`silly nephew, there cannot be %d cows`, cows))
	} else {
		return 0, errors.New("non-scale error")
	}
}
*/
package main

func main() {

}
