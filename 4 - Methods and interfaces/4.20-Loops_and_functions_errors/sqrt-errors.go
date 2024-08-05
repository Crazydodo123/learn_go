package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negatime number: %s", fmt.Sprint(float64(e)))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for math.Abs(z*z-x) > 0.0000001 {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	ans, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}

	ans, err = Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
