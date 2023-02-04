package main

import (
	"encoding/json"
	"fmt"
)

type Tree struct {
	Value float32 `json:"value"`
	Left  *Tree   `json:"left,omitempty"`
	Right *Tree   `json:"right,omitempty"`
}

func main() {
	rawValue := `{ "value": 100.0, "left": { "value": 180.0 } }`
	value := Tree{}
	if err := json.Unmarshal([]byte(rawValue), &value); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", value)

}
