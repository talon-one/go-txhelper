// Code generated by 'yaegi extract encoding/json'. DO NOT EDIT.

// +build go1.16

package stdlib

import (
	"encoding/json"
	"reflect"
)

func init() {
	Symbols["encoding/json/json"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Compact":       reflect.ValueOf(json.Compact),
		"HTMLEscape":    reflect.ValueOf(json.HTMLEscape),
		"Indent":        reflect.ValueOf(json.Indent),
		"Marshal":       reflect.ValueOf(json.Marshal),
		"MarshalIndent": reflect.ValueOf(json.MarshalIndent),
		"NewDecoder":    reflect.ValueOf(json.NewDecoder),
		"NewEncoder":    reflect.ValueOf(json.NewEncoder),
		"Unmarshal":     reflect.ValueOf(json.Unmarshal),
		"Valid":         reflect.ValueOf(json.Valid),

		// type definitions
		"Decoder":               reflect.ValueOf((*json.Decoder)(nil)),
		"Delim":                 reflect.ValueOf((*json.Delim)(nil)),
		"Encoder":               reflect.ValueOf((*json.Encoder)(nil)),
		"InvalidUTF8Error":      reflect.ValueOf((*json.InvalidUTF8Error)(nil)),
		"InvalidUnmarshalError": reflect.ValueOf((*json.InvalidUnmarshalError)(nil)),
		"Marshaler":             reflect.ValueOf((*json.Marshaler)(nil)),
		"MarshalerError":        reflect.ValueOf((*json.MarshalerError)(nil)),
		"Number":                reflect.ValueOf((*json.Number)(nil)),
		"RawMessage":            reflect.ValueOf((*json.RawMessage)(nil)),
		"SyntaxError":           reflect.ValueOf((*json.SyntaxError)(nil)),
		"Token":                 reflect.ValueOf((*json.Token)(nil)),
		"UnmarshalFieldError":   reflect.ValueOf((*json.UnmarshalFieldError)(nil)),
		"UnmarshalTypeError":    reflect.ValueOf((*json.UnmarshalTypeError)(nil)),
		"Unmarshaler":           reflect.ValueOf((*json.Unmarshaler)(nil)),
		"UnsupportedTypeError":  reflect.ValueOf((*json.UnsupportedTypeError)(nil)),
		"UnsupportedValueError": reflect.ValueOf((*json.UnsupportedValueError)(nil)),

		// interface wrapper definitions
		"_Marshaler":   reflect.ValueOf((*_encoding_json_Marshaler)(nil)),
		"_Token":       reflect.ValueOf((*_encoding_json_Token)(nil)),
		"_Unmarshaler": reflect.ValueOf((*_encoding_json_Unmarshaler)(nil)),
	}
}

// _encoding_json_Marshaler is an interface wrapper for Marshaler type
type _encoding_json_Marshaler struct {
	IValue       interface{}
	WMarshalJSON func() ([]byte, error)
}

func (W _encoding_json_Marshaler) MarshalJSON() ([]byte, error) { return W.WMarshalJSON() }

// _encoding_json_Token is an interface wrapper for Token type
type _encoding_json_Token struct {
	IValue interface{}
}

// _encoding_json_Unmarshaler is an interface wrapper for Unmarshaler type
type _encoding_json_Unmarshaler struct {
	IValue         interface{}
	WUnmarshalJSON func(a0 []byte) error
}

func (W _encoding_json_Unmarshaler) UnmarshalJSON(a0 []byte) error { return W.WUnmarshalJSON(a0) }
