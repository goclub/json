package xjson

import (
)

var emptyObjectAndArrayNotNull = true
func emptyObjectString () string {
	if emptyObjectAndArrayNotNull {
		return "{}"
	} else {
		return "null"
	}
}
func emptyArrayString () string {
	if emptyObjectAndArrayNotNull {
		return "[]"
	} else {
		return "null"
	}
}
var autoNumberConvertString = true

