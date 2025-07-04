package fullmath

import (
	cons "github.com/otbzi/uniswap-simulator/lib/constants"
	ui "github.com/otbzi/uniswap-simulator/uint256"
)

func MulDivRoundingUp(a, b, denominator *ui.Int) *ui.Int {
	if a.IsZero() || b.IsZero() {
		return ui.NewInt(0)
	}
	product := ui.Umul(a, b)
	var q [5]uint64
	quotient := q[:]
	rem := ui.Udivrem(quotient, product[:], denominator)
	result := (*ui.Int)(quotient[0:4])
	if !rem.IsZero() {
		result.Add(result, cons.One)
	}
	return result
}

func MulDiv(a, b, denominator *ui.Int) *ui.Int {
	if a.IsZero() || b.IsZero() {
		return ui.NewInt(0)
	}
	product := ui.Umul(a, b)
	var q [5]uint64
	quotient := q[:]
	ui.Udivrem(quotient, product[:], denominator)
	result := (*ui.Int)(quotient[0:4])
	return result
}
