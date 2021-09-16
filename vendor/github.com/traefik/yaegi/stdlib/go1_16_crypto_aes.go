// Code generated by 'yaegi extract crypto/aes'. DO NOT EDIT.

// +build go1.16,!go1.17

package stdlib

import (
	"crypto/aes"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["crypto/aes/aes"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BlockSize": reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"NewCipher": reflect.ValueOf(aes.NewCipher),

		// type definitions
		"KeySizeError": reflect.ValueOf((*aes.KeySizeError)(nil)),
	}
}
