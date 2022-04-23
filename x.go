package xjson

import "log"

var emptyObjectAndArrayNotNull = true

func emptyObjectString() string {
	if emptyObjectAndArrayNotNull {
		return "{}"
	} else {
		return "null"
	}
}
func emptyArrayString() string {
	if emptyObjectAndArrayNotNull {
		return "[]"
	} else {
		return "null"
	}
}

var autoNumberConvertString = true

func Print(name string, v interface{}) {
	data, err := Marshal(v)
	if err != nil {
		log.Print(name, "\n", err)
	} else {
		log.Print(name, "\n", string(data))
	}
}
func PrintIndent(name string, v interface{}) {
	data, err := MarshalIndent(v, "", "  ")
	if err != nil {
		log.Print(name, "\n", err)
	} else {
		log.Print(name, "\n", string(data))
	}
}
