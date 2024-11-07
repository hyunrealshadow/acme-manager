package util

import (
	"fmt"
	"reflect"
	"unsafe"
)

func GetStructPtrUnExportedField(source interface{}, fieldName string) reflect.Value {
	// Get the reflect.Value of the unexported field
	v := reflect.ValueOf(source).Elem().FieldByName(fieldName)
	// Create an addressable reflect.Value pointing to the field
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

func SetStructPtrUnExportedStrField(source interface{}, fieldName string, fieldVal interface{}) (err error) {
	v := GetStructPtrUnExportedField(source, fieldName)
	rv := reflect.ValueOf(fieldVal)
	if v.Kind() != rv.Kind() {
		return fmt.Errorf("invalid kind: expected kind %v, got kind: %v", v.Kind(), rv.Kind())
	}
	// Set the value of the unexported field
	v.Set(rv)
	return nil
}
