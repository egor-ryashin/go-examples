package main

import (
	"encoding/json"
	"fmt"
)

func json_test() {
	type Rectangle struct {
		Width  float64
		Height float64
	}
	var rect Rectangle
	json.Unmarshal([]byte("{\"width\": 28}"), &rect)
	fmt.Println(rect.Width)
}
