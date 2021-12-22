// Code generated by 'yaegi extract crypto/x509/pkix'. DO NOT EDIT.

// +build go1.16,!go1.17

package stdlib

import (
	"crypto/x509/pkix"
	"reflect"
)

func init() {
	Symbols["crypto/x509/pkix/pkix"] = map[string]reflect.Value{
		// type definitions
		"AlgorithmIdentifier":          reflect.ValueOf((*pkix.AlgorithmIdentifier)(nil)),
		"AttributeTypeAndValue":        reflect.ValueOf((*pkix.AttributeTypeAndValue)(nil)),
		"AttributeTypeAndValueSET":     reflect.ValueOf((*pkix.AttributeTypeAndValueSET)(nil)),
		"CertificateList":              reflect.ValueOf((*pkix.CertificateList)(nil)),
		"Extension":                    reflect.ValueOf((*pkix.Extension)(nil)),
		"Name":                         reflect.ValueOf((*pkix.Name)(nil)),
		"RDNSequence":                  reflect.ValueOf((*pkix.RDNSequence)(nil)),
		"RelativeDistinguishedNameSET": reflect.ValueOf((*pkix.RelativeDistinguishedNameSET)(nil)),
		"RevokedCertificate":           reflect.ValueOf((*pkix.RevokedCertificate)(nil)),
		"TBSCertificateList":           reflect.ValueOf((*pkix.TBSCertificateList)(nil)),
	}
}
