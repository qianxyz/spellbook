package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	for _, spell := range Spells {
		bytes, _ := json.MarshalIndent(spell, "", "  ")
		fmt.Println(string(bytes))
	}
}
