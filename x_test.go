package xjson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init () {
	emptyObjectAndArrayNotNull = false
	autoNumberConvertString = false
}


func TestNilSliceConvertEmptyArray(t *testing.T) {
	emptyObjectAndArrayNotNull = true
	defer func() {
		emptyObjectAndArrayNotNull = false
	}()
	a := struct {
		Bools []bool
		Ints []int
		Structs []struct{A int}
		Bytes []byte
		Strings []string
		Map map[int]int
	}{}
	data, err := Marshal(a) ; if err != nil {panic(err)}
	actual := string(data)
	expect := `{"Bools":[],"Ints":[],"Structs":[],"Bytes":[],"Strings":[],"Map":{}}`
	assert.Equal(t,expect, actual)
}

func TestUnmarshalIntConvertString(t *testing.T) {
	autoNumberConvertString = true
	defer func() {
		autoNumberConvertString = false
	}()
	data := []byte(`{"Int":1,"Int8":2,"Int16":3,"Int32":4,"Int64":5,"Uint":1,"Uint8":2,"Uint16":3,"Uint32":4,"Uint64":5,"Float32":1.111, "Float64":2.222}`)
	query := struct {
		Int     string
		Int8    string
		Int16   string
		Int32   string
		Int64   string
		Uint    string
		Uint8   string
		Uint16  string
		Uint32  string
		Uint64  string
		Float32 string
		Float64 string
	}{}
	err := Unmarshal(data, &query)
	assert.NoError(t, err)
	assert.Equal(t,"1", query.Int)
	assert.Equal(t,"2", query.Int8)
	assert.Equal(t,"3", query.Int16)
	assert.Equal(t,"4", query.Int32)
	assert.Equal(t,"5", query.Int64)
	assert.Equal(t,"1", query.Uint)
	assert.Equal(t,"2", query.Uint8)
	assert.Equal(t,"3", query.Uint16)
	assert.Equal(t,"4", query.Uint32)
	assert.Equal(t,"5", query.Uint64)
	assert.Equal(t,"1.111", query.Float32)
	assert.Equal(t,"2.222", query.Float64)
}

func TestUnmarshalStringConvertInt (t *testing.T) {
	autoNumberConvertString = true
	defer func() {
		autoNumberConvertString = false
	}()
	query := struct {
		Int     int
		Int8    int8
		Int16   int16
		Int32   int32
		Int64   int64
		Uint    uint
		Uint8   uint8
		Uint16  uint16
		Uint32  uint32
		Uint64  uint64
		Float32 float32
		Float64 float64
	}{}
	err := Unmarshal([]byte(`{"Int":"1","Int8":"2","Int16":"3","Int32":"4","Int64":"5","Uint":"1","Uint8":"2","Uint16":"3","Uint32":"4","Uint64":"5","Float32": "1.111", "Float64": "2.222"}`),&query)
	assert.NoError(t, err)
	assert.Equal(t,1, query.Int)
	assert.Equal(t,int8(2), query.Int8)
	assert.Equal(t,int16(3), query.Int16)
	assert.Equal(t,int32(4), query.Int32)
	assert.Equal(t,int64(5), query.Int64)
	assert.Equal(t,uint(1), query.Uint)
	assert.Equal(t,uint8(2), query.Uint8)
	assert.Equal(t,uint16(3), query.Uint16)
	assert.Equal(t,uint32(4), query.Uint32)
	assert.Equal(t,uint64(5), query.Uint64)
	assert.Equal(t,float32(1.111), query.Float32)
	assert.Equal(t,float64(2.222), query.Float64)

	{
		a := struct {
			A int
		}{}
		err := Unmarshal([]byte(`{"A":"xx"}`), &a)
		assert.EqualError(t, err, "json: cannot unmarshal string into Go struct field .A of type int")
	}
}

func TestPrint(t *testing.T) {
	v := struct {
		Name string `json:"name"`
		Age uint `json:"age"`
	}{Name: "nimoc", Age: 18}
	Print("v", v)
}
