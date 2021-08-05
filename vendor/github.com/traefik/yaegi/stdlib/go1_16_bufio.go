// Code generated by 'yaegi extract bufio'. DO NOT EDIT.

// +build go1.16

package stdlib

import (
	"bufio"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["bufio/bufio"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrAdvanceTooFar":     reflect.ValueOf(&bufio.ErrAdvanceTooFar).Elem(),
		"ErrBadReadCount":      reflect.ValueOf(&bufio.ErrBadReadCount).Elem(),
		"ErrBufferFull":        reflect.ValueOf(&bufio.ErrBufferFull).Elem(),
		"ErrFinalToken":        reflect.ValueOf(&bufio.ErrFinalToken).Elem(),
		"ErrInvalidUnreadByte": reflect.ValueOf(&bufio.ErrInvalidUnreadByte).Elem(),
		"ErrInvalidUnreadRune": reflect.ValueOf(&bufio.ErrInvalidUnreadRune).Elem(),
		"ErrNegativeAdvance":   reflect.ValueOf(&bufio.ErrNegativeAdvance).Elem(),
		"ErrNegativeCount":     reflect.ValueOf(&bufio.ErrNegativeCount).Elem(),
		"ErrTooLong":           reflect.ValueOf(&bufio.ErrTooLong).Elem(),
		"MaxScanTokenSize":     reflect.ValueOf(constant.MakeFromLiteral("65536", token.INT, 0)),
		"NewReadWriter":        reflect.ValueOf(bufio.NewReadWriter),
		"NewReader":            reflect.ValueOf(bufio.NewReader),
		"NewReaderSize":        reflect.ValueOf(bufio.NewReaderSize),
		"NewScanner":           reflect.ValueOf(bufio.NewScanner),
		"NewWriter":            reflect.ValueOf(bufio.NewWriter),
		"NewWriterSize":        reflect.ValueOf(bufio.NewWriterSize),
		"ScanBytes":            reflect.ValueOf(bufio.ScanBytes),
		"ScanLines":            reflect.ValueOf(bufio.ScanLines),
		"ScanRunes":            reflect.ValueOf(bufio.ScanRunes),
		"ScanWords":            reflect.ValueOf(bufio.ScanWords),

		// type definitions
		"ReadWriter": reflect.ValueOf((*bufio.ReadWriter)(nil)),
		"Reader":     reflect.ValueOf((*bufio.Reader)(nil)),
		"Scanner":    reflect.ValueOf((*bufio.Scanner)(nil)),
		"SplitFunc":  reflect.ValueOf((*bufio.SplitFunc)(nil)),
		"Writer":     reflect.ValueOf((*bufio.Writer)(nil)),
	}
}
