package utils

import (
	"encoding/json"
	"reflect"
)

func StructToMapViaJson(object interface{}) (m map[string]interface{}) {
	j, _ := json.Marshal(object)
	json.Unmarshal(j, &m)
	return m
}

func StructToMapViaReflect(object interface{}) (m map[string]interface{}) {

	if m == nil {
		m = make(map[string]interface{})
	}

	elem := reflect.ValueOf(object).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		m[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return m
}
