package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Resp1 struct {
	Page   int
	Fruits []string
}

type Resp2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1 := &Resp1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res1B, _ := json.Marshal(res1)
	fmt.Println(string(res1B))

	res2 := &Resp2{
		Page:   2,
		Fruits: []string{"apple", "peach", "pear"}}

	res2B, _ := json.Marshal(res2)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(data["num"])
	fmt.Printf("%T\n", data["num"])
	fmt.Printf("%T\n", data["num"].(float64))

	fmt.Println(data["strs"])
	fmt.Printf("%T\n", data["strs"])
	str1 := data["strs"].([]interface{})
	fmt.Printf("%T\n", str1[0])

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Resp2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
