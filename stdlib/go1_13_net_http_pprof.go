// Code generated by 'goexports net/http/pprof'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"net/http/pprof"
	"reflect"
)

func init() {
	Symbols["net/http/pprof"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Cmdline": reflect.ValueOf(pprof.Cmdline),
		"Handler": reflect.ValueOf(pprof.Handler),
		"Index":   reflect.ValueOf(pprof.Index),
		"Profile": reflect.ValueOf(pprof.Profile),
		"Symbol":  reflect.ValueOf(pprof.Symbol),
		"Trace":   reflect.ValueOf(pprof.Trace),
	}
}
