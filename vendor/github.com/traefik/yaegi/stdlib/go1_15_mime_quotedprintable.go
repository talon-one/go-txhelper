// Code generated by 'yaegi extract mime/quotedprintable'. DO NOT EDIT.

// +build go1.15,!go1.16

package stdlib

import (
	"mime/quotedprintable"
	"reflect"
)

func init() {
	Symbols["mime/quotedprintable"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewReader": reflect.ValueOf(quotedprintable.NewReader),
		"NewWriter": reflect.ValueOf(quotedprintable.NewWriter),

		// type definitions
		"Reader": reflect.ValueOf((*quotedprintable.Reader)(nil)),
		"Writer": reflect.ValueOf((*quotedprintable.Writer)(nil)),
	}
}