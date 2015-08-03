package main

import (
	"encoding/json"
	"fmt"
)

/**
参考
http://blog.golang.org/json-and-go
https://eager.io/blog/go-and-json/
 */

func main() {
	//testMarshal()
	//testUnmarshal()
	//testType()
	testArbitraryData()
}

/**
The json package only accesses the exported fields of struct types (those that begin with an uppercase letter).
Therefore only the the exported fields of a struct will be present in the JSON output.
 */
type Message struct {
	Name string
	Body string
	Time int64
}

func testMarshal() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)

	if err!=nil {
		fmt.Println("error in marshal")
	}
	fmt.Println(string(b))

}

func testUnmarshal() {
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	var m Message
	err := json.Unmarshal(b, &m)
	if err!=nil {
		fmt.Println("error in unmarshal")
	}
	fmt.Println(m)
}
/**
The json package uses map[string]interface{} and []interface{} values to store arbitrary JSON objects and arrays;
it will happily unmarshal any valid JSON blob into a plain interface{} value.
The default concrete Go types are:
	bool for JSON booleans,
	float64 for JSON numbers,
	string for JSON strings, and
	nil for JSON null.
 */
func testArbitraryData() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err!=nil {
		fmt.Println("error in unmarshal")
	}
	m := f.(map[string]interface{})
	fmt.Println(m)


	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}




}

func testReferenceType() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	type FamilyMember struct {
		Name    string
		Age     int
		Parents []string
	}

	var fm FamilyMember
	err1 := json.Unmarshal(b, &fm)
	if err1!=nil {
		fmt.Println(err1)
	}
	fmt.Println(fm)
}




func testType() {
	var i interface{}
	i = "a string"
	i = 2011
	i = 2.777


	switch  v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case float64:
		fmt.Println("float64", v)
	}

}
