package main

import (
	"encoding/json"
	"fmt"
)

type Jsonstructor struct {
	beer string
	tea  string
}

func main() {
	f := Jsonstructor{"royal", "mcdonald"}
	b, _ := json.Marshal(f)
	fmt.Println(string(b), f)
	json.Unmarshal(b, &f)
	fmt.Println(string(b), &f)
}
