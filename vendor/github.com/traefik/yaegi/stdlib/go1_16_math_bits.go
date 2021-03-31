// Code generated by 'yaegi extract math/bits'. DO NOT EDIT.

// +build go1.16

package stdlib

import (
	"go/constant"
	"go/token"
	"math/bits"
	"reflect"
)

func init() {
	Symbols["math/bits"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Add":             reflect.ValueOf(bits.Add),
		"Add32":           reflect.ValueOf(bits.Add32),
		"Add64":           reflect.ValueOf(bits.Add64),
		"Div":             reflect.ValueOf(bits.Div),
		"Div32":           reflect.ValueOf(bits.Div32),
		"Div64":           reflect.ValueOf(bits.Div64),
		"LeadingZeros":    reflect.ValueOf(bits.LeadingZeros),
		"LeadingZeros16":  reflect.ValueOf(bits.LeadingZeros16),
		"LeadingZeros32":  reflect.ValueOf(bits.LeadingZeros32),
		"LeadingZeros64":  reflect.ValueOf(bits.LeadingZeros64),
		"LeadingZeros8":   reflect.ValueOf(bits.LeadingZeros8),
		"Len":             reflect.ValueOf(bits.Len),
		"Len16":           reflect.ValueOf(bits.Len16),
		"Len32":           reflect.ValueOf(bits.Len32),
		"Len64":           reflect.ValueOf(bits.Len64),
		"Len8":            reflect.ValueOf(bits.Len8),
		"Mul":             reflect.ValueOf(bits.Mul),
		"Mul32":           reflect.ValueOf(bits.Mul32),
		"Mul64":           reflect.ValueOf(bits.Mul64),
		"OnesCount":       reflect.ValueOf(bits.OnesCount),
		"OnesCount16":     reflect.ValueOf(bits.OnesCount16),
		"OnesCount32":     reflect.ValueOf(bits.OnesCount32),
		"OnesCount64":     reflect.ValueOf(bits.OnesCount64),
		"OnesCount8":      reflect.ValueOf(bits.OnesCount8),
		"Rem":             reflect.ValueOf(bits.Rem),
		"Rem32":           reflect.ValueOf(bits.Rem32),
		"Rem64":           reflect.ValueOf(bits.Rem64),
		"Reverse":         reflect.ValueOf(bits.Reverse),
		"Reverse16":       reflect.ValueOf(bits.Reverse16),
		"Reverse32":       reflect.ValueOf(bits.Reverse32),
		"Reverse64":       reflect.ValueOf(bits.Reverse64),
		"Reverse8":        reflect.ValueOf(bits.Reverse8),
		"ReverseBytes":    reflect.ValueOf(bits.ReverseBytes),
		"ReverseBytes16":  reflect.ValueOf(bits.ReverseBytes16),
		"ReverseBytes32":  reflect.ValueOf(bits.ReverseBytes32),
		"ReverseBytes64":  reflect.ValueOf(bits.ReverseBytes64),
		"RotateLeft":      reflect.ValueOf(bits.RotateLeft),
		"RotateLeft16":    reflect.ValueOf(bits.RotateLeft16),
		"RotateLeft32":    reflect.ValueOf(bits.RotateLeft32),
		"RotateLeft64":    reflect.ValueOf(bits.RotateLeft64),
		"RotateLeft8":     reflect.ValueOf(bits.RotateLeft8),
		"Sub":             reflect.ValueOf(bits.Sub),
		"Sub32":           reflect.ValueOf(bits.Sub32),
		"Sub64":           reflect.ValueOf(bits.Sub64),
		"TrailingZeros":   reflect.ValueOf(bits.TrailingZeros),
		"TrailingZeros16": reflect.ValueOf(bits.TrailingZeros16),
		"TrailingZeros32": reflect.ValueOf(bits.TrailingZeros32),
		"TrailingZeros64": reflect.ValueOf(bits.TrailingZeros64),
		"TrailingZeros8":  reflect.ValueOf(bits.TrailingZeros8),
		"UintSize":        reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
	}
}
