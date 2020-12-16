# xjson

```go
import "github.com/goclub/json"
```
Go JSON standard package extension library.

> Go 官方 JSON 标准库扩展

## Extension feature

> 扩展特性

1. nil slice nad nil map marshal to `[]` or `{}` instead of null.
> 1. 空切片空 map 序列化字符串为 `[]` 或 `{}` 而不是 `null`。


```go
response := struct {
    Books []string `json:"books"`
    Map map[string]string `json:"map"`
}{}
data, err := xjson.Marshal(response) ; if err != nil {panic(err)}
log.Print(string(data)) // {"books":[], map:{}}
```

2. auto int and float convert to stirng instead return error.

> 自动int和 float到字符串，而不是返回错误

```go
request := struct {
    Page int `json:"page"`
    Price float64 `json:"price"`
}{}
err := xjson.Unmarshal([]byte(`{"page":"1", "price": "1.05"}`), &request) ; if err != nil {panic(err)}
log.Printf("%+v", request) // {Page:1 Price:1.05}
```

```go
request := struct {
   ID string `json:"id"`
}{}
err := xjson.Unmarshal([]byte(`{"id":1`), &request) ; if err != nil {panic(err)}
log.Printf("%+v", request) // {ID:"1"}
```
