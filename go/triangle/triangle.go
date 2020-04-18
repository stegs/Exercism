// Package triangle contains the code to determine the kind of triangle
package triangle

import (
	"math"
)

type Kind string

const (
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// KindFromSides takes three sides of a triangle and returns the Kind of triangle it is.
func KindFromSides(a, b, c float64) Kind {

	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if (a == b && a != c) || (a == c && a != b) || (a != b && b == c) {
		return Iso
	}

	if a != b && a != c {
		return Sca
	}

	return NaT
}
