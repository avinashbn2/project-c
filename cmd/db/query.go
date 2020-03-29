package db

import (
	"fmt"
	"reflect"
)

// DBFields reflects on a struct and returns the values of fields with `db` tags,
// or a map[string]interface{} and returns the keys.
func DBFields(values interface{}) []string {
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	fields := []string{}
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("json")
			if field != "" {
				fields = append(fields, field)
			}
		}
		return fields
	}
	if v.Kind() == reflect.Map {
		for _, keyv := range v.MapKeys() {
			fields = append(fields, keyv.String())
		}
		return fields
	}
	panic(fmt.Errorf("DBFields requires a struct or a map, found: %s", v.Kind().String()))
}
