// Code generated by 'yaegi extract image/png'. DO NOT EDIT.

// +build go1.16

package stdlib

import (
	"image/png"
	"reflect"
)

func init() {
	Symbols["image/png/png"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BestCompression":    reflect.ValueOf(png.BestCompression),
		"BestSpeed":          reflect.ValueOf(png.BestSpeed),
		"Decode":             reflect.ValueOf(png.Decode),
		"DecodeConfig":       reflect.ValueOf(png.DecodeConfig),
		"DefaultCompression": reflect.ValueOf(png.DefaultCompression),
		"Encode":             reflect.ValueOf(png.Encode),
		"NoCompression":      reflect.ValueOf(png.NoCompression),

		// type definitions
		"CompressionLevel":  reflect.ValueOf((*png.CompressionLevel)(nil)),
		"Encoder":           reflect.ValueOf((*png.Encoder)(nil)),
		"EncoderBuffer":     reflect.ValueOf((*png.EncoderBuffer)(nil)),
		"EncoderBufferPool": reflect.ValueOf((*png.EncoderBufferPool)(nil)),
		"FormatError":       reflect.ValueOf((*png.FormatError)(nil)),
		"UnsupportedError":  reflect.ValueOf((*png.UnsupportedError)(nil)),

		// interface wrapper definitions
		"_EncoderBufferPool": reflect.ValueOf((*_image_png_EncoderBufferPool)(nil)),
	}
}

// _image_png_EncoderBufferPool is an interface wrapper for EncoderBufferPool type
type _image_png_EncoderBufferPool struct {
	WGet func() *png.EncoderBuffer
	WPut func(a0 *png.EncoderBuffer)
}

func (W _image_png_EncoderBufferPool) Get() *png.EncoderBuffer   { return W.WGet() }
func (W _image_png_EncoderBufferPool) Put(a0 *png.EncoderBuffer) { W.WPut(a0) }
