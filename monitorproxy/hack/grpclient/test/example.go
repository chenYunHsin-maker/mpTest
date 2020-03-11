package main

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	json "github.com/golang/protobuf/jsonpb"

	"reflect"
	"strings"
)

func encodeElem(v reflect.Value) {
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			vf := v.Field(i)
			if vf.Kind() == reflect.Slice {
				vf.Set(reflect.MakeSlice(vf.Type(), 1, 1))
				vf.SetLen(1)
				if strings.Contains(vf.Type().String(), "*") {
					vf.Index(0).Set(reflect.New(vf.Index(0).Type().Elem()))
					encodeElem(vf.Index(0).Elem())
				} else {
					encodeElem(vf.Index(0))
				}
			} else if vf.Kind() == reflect.Struct {
				encodeElem(vf)
			} else if vf.Kind() == reflect.Ptr {
				vf.Set(reflect.New(vf.Type().Elem()))
				encodeElem(vf.Elem())
			}
		}
	}
}

func exampleJSON(v interface{}, pb proto.Message) string {
	vv := reflect.ValueOf(v)
	ve := vv.Elem()
	encodeElem(ve)

	m := &json.Marshaler{EmitDefaults: true, Indent: "   ", OrigName: true}
	js, _ := m.MarshalToString(pb)
	fmt.Println(js)
	fmt.Println()
	m = &json.Marshaler{EmitDefaults: true, OrigName: true}
	js, _ = m.MarshalToString(pb)
	return fmt.Sprintf("\"%s\"", strings.Replace(string(js), "\"", "'", -1))
}
