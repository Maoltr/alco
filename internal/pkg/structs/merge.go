package structs

import (
	"reflect"
)

// Merge receives two structs, and merges them excluding fields with tag name: `structs`, value "-"
func Merge(dst, src interface{}) {
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)
	if srcValue.Kind() != reflect.Ptr || dstValue.Kind() != reflect.Ptr {
		return
	}

	for i := 0; i < srcValue.Elem().NumField(); i++ {
		newValue := srcValue.Elem().Field(i)
		fieldName := srcValue.Elem().Type().Field(i).Name

		skip := srcValue.Elem().Type().Field(i).Tag.Get("structs")
		if skip == "-" {
			continue
		}

		if newValue.Kind() > reflect.Float64 &&
			newValue.Kind() != reflect.String &&
			newValue.Kind() != reflect.Struct &&
			newValue.Kind() != reflect.Ptr &&
			newValue.Kind() != reflect.Slice {
			continue
		}

		if newValue.Kind() != reflect.Ptr {
			continue
		}

		// Field is pointer check if it's nil or set
		if newValue.IsNil() {
			continue
		}

		// Field is set assign it to dest
		if dstValue.Elem().FieldByName(fieldName).Kind() == reflect.Ptr {
			dstValue.Elem().FieldByName(fieldName).Set(newValue)
			continue
		}

		oldValue := dstValue.Elem().FieldByName(fieldName)
		if oldValue.IsValid() {
			oldValue.Set(newValue.Elem())
		}
	}
}
