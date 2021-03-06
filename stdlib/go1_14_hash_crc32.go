// Code generated by 'goexports hash/crc32'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"hash/crc32"
	"reflect"
)

func init() {
	Symbols["hash/crc32"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Castagnoli":   reflect.ValueOf(uint32(crc32.Castagnoli)),
		"Checksum":     reflect.ValueOf(crc32.Checksum),
		"ChecksumIEEE": reflect.ValueOf(crc32.ChecksumIEEE),
		"IEEE":         reflect.ValueOf(uint32(crc32.IEEE)),
		"IEEETable":    reflect.ValueOf(&crc32.IEEETable).Elem(),
		"Koopman":      reflect.ValueOf(uint32(crc32.Koopman)),
		"MakeTable":    reflect.ValueOf(crc32.MakeTable),
		"New":          reflect.ValueOf(crc32.New),
		"NewIEEE":      reflect.ValueOf(crc32.NewIEEE),
		"Size":         reflect.ValueOf(crc32.Size),
		"Update":       reflect.ValueOf(crc32.Update),

		// type definitions
		"Table": reflect.ValueOf((*crc32.Table)(nil)),
	}
}
