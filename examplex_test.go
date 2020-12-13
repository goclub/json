package xjson

import (
	"log"
	"testing"
)

func ExampleEmptyArrayAndEmptyObject () {
	response := struct {
		Books []string `json:"books"`
		Map map[string]string `json:"map"`
	}{}
	data, err := Marshal(response) ; if err != nil {panic(err)}
	log.Print(string(data)) // {"books":null, map:{}}
}
func TestExampleEmptyArrayAndEmptyObject(t *testing.T) {
	emptyObjectAndArrayNotNull = true
	defer func() {
		emptyObjectAndArrayNotNull = false
	}()
	ExampleEmptyArrayAndEmptyObject()
}
func ExampleIntFloatAutoConvertString () {
	request := struct {
		Page int `json:"page"`
		Price float64 `json:"price"`
	}{}
	err := Unmarshal([]byte(`{"page":"1", "price": "1.05"}`), &request) ; if err != nil {panic(err)}
	log.Printf("%+v", request) // {Page:1 Price:1.05}
}
func TestExampleIntFloatAutoConvertString(t *testing.T) {
	autoNumberConvertString = true
	defer func() {
		autoNumberConvertString = false
	}()
	ExampleIntFloatAutoConvertString()
}
