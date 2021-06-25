// Code generated by 'yaegi extract sort'. DO NOT EDIT.

// +build go1.15,!go1.16

package stdlib

import (
	"reflect"
	"sort"
)

func init() {
	Symbols["sort/sort"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Float64s":          reflect.ValueOf(sort.Float64s),
		"Float64sAreSorted": reflect.ValueOf(sort.Float64sAreSorted),
		"Ints":              reflect.ValueOf(sort.Ints),
		"IntsAreSorted":     reflect.ValueOf(sort.IntsAreSorted),
		"IsSorted":          reflect.ValueOf(sort.IsSorted),
		"Reverse":           reflect.ValueOf(sort.Reverse),
		"Search":            reflect.ValueOf(sort.Search),
		"SearchFloat64s":    reflect.ValueOf(sort.SearchFloat64s),
		"SearchInts":        reflect.ValueOf(sort.SearchInts),
		"SearchStrings":     reflect.ValueOf(sort.SearchStrings),
		"Slice":             reflect.ValueOf(sort.Slice),
		"SliceIsSorted":     reflect.ValueOf(sort.SliceIsSorted),
		"SliceStable":       reflect.ValueOf(sort.SliceStable),
		"Sort":              reflect.ValueOf(sort.Sort),
		"Stable":            reflect.ValueOf(sort.Stable),
		"Strings":           reflect.ValueOf(sort.Strings),
		"StringsAreSorted":  reflect.ValueOf(sort.StringsAreSorted),

		// type definitions
		"Float64Slice": reflect.ValueOf((*sort.Float64Slice)(nil)),
		"IntSlice":     reflect.ValueOf((*sort.IntSlice)(nil)),
		"Interface":    reflect.ValueOf((*sort.Interface)(nil)),
		"StringSlice":  reflect.ValueOf((*sort.StringSlice)(nil)),

		// interface wrapper definitions
		"_Interface": reflect.ValueOf((*_sort_Interface)(nil)),
	}
}

// _sort_Interface is an interface wrapper for Interface type
type _sort_Interface struct {
	WLen  func() int
	WLess func(i int, j int) bool
	WSwap func(i int, j int)
}

func (W _sort_Interface) Len() int               { return W.WLen() }
func (W _sort_Interface) Less(i int, j int) bool { return W.WLess(i, j) }
func (W _sort_Interface) Swap(i int, j int)      { W.WSwap(i, j) }
