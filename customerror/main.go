package main

import (
	"errors"
	"fmt"
)

func volume(r float64) (float64, error) {

	if r < 0 {

		return 0, errors.New("Volume calculation failed: radius negative.")
	}
	return (4.0 / 3.0) * 3.14 * r * r * r
}

func main() {

	radius := 1.0
	vol := volume(radius)
	fmt.Printf("volume of shpere is %0.2f\n", vol)

}
